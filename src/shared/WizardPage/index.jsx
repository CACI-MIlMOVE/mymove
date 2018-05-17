import React, { Component } from 'react'; // eslint-disable-line
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import windowSize from 'react-window-size';
import { bindActionCreators } from 'redux';
import { withRouter } from 'react-router-dom';
import { push } from 'react-router-redux';
import Alert from 'shared/Alert'; // eslint-disable-line
import generatePath from './generatePath';
import './index.css';
import { mobileSize } from 'shared/constants';

import {
  getNextPagePath,
  getPreviousPagePath,
  isFirstPage,
  isLastPage,
} from './utils';

export class WizardPage extends Component {
  constructor(props) {
    super(props);
    this.nextPage = this.nextPage.bind(this);
    this.previousPage = this.previousPage.bind(this);
    this.cancelFlow = this.cancelFlow.bind(this);
    this.state = { transitionFunc: null };
  }
  componentDidUpdate() {
    if (this.props.hasSucceeded) this.onSubmitSuccessful();
    if (this.props.error) window.scrollTo(0, 0);
  }
  componentDidMount() {
    window.scrollTo(0, 0);
  }
  cancelFlow() {
    this.props.push(`/`);
  }
  beforeTransition(func) {
    const {
      isAsync,
      pageIsDirty,
      pageList,
      pageKey,
      handleSubmit,
    } = this.props;
    const path = func(pageList, pageKey);
    if (pageIsDirty && handleSubmit) {
      handleSubmit();
      if (isAsync) {
        this.setState({ transitionFunc: func });
      } else {
        this.goto(path);
      }
    } else {
      this.goto(path);
    }
  }
  goto(path) {
    const {
      push,
      match: { params },
      additionalParams,
    } = this.props;
    const combinedParams = additionalParams
      ? Object.assign({}, additionalParams, params)
      : params;
    // comes from react router redux: doing this moves to the route at path  (might consider going back to history since we need withRouter)
    push(generatePath(path, combinedParams));
  }
  onSubmitSuccessful() {
    const { transitionFunc } = this.state;
    const { pageKey, pageList } = this.props;
    if (transitionFunc) this.goto(transitionFunc(pageList, pageKey));
  }
  nextPage() {
    this.beforeTransition(getNextPagePath);
  }

  previousPage() {
    this.beforeTransition(getPreviousPagePath);
  }

  render() {
    const isMobile = this.props.windowWidth < mobileSize;
    const {
      handleSubmit,
      pageKey,
      pageList,
      children,
      error,
      pageIsValid,
      pageIsDirty,
    } = this.props;
    const canMoveForward = pageIsValid;
    const canMoveBackward =
      (pageIsValid || !pageIsDirty) && !isFirstPage(pageList, pageKey);
    return (
      <div className="usa-grid">
        {error && (
          <div className="usa-width-one-whole error-message">
            <Alert type="error" heading="An error occurred">
              {error.message}
            </Alert>
          </div>
        )}
        <div className="usa-width-one-whole">{children}</div>
        <div className="usa-width-one-whole lower-nav-btns">
          {!isMobile && (
            <div className="left cancel">
              <button
                className="usa-button-secondary"
                onClick={this.cancelFlow}
              >
                Cancel
              </button>
            </div>
          )}
          <div className="prev-next">
            <button
              className="usa-button-secondary prev"
              onClick={this.previousPage}
              disabled={!canMoveBackward}
            >
              Back
            </button>
            {!isLastPage(pageList, pageKey) && (
              <button
                className="usa-button-primary next"
                onClick={this.nextPage}
                disabled={!canMoveForward}
              >
                Next
              </button>
            )}
            {isLastPage(pageList, pageKey) && (
              <button
                className="usa-button-primary next"
                onClick={handleSubmit}
                disabled={!canMoveForward}
              >
                Complete
              </button>
            )}
          </div>
        </div>
      </div>
    );
  }
}

WizardPage.propTypes = {
  handleSubmit: PropTypes.func.isRequired,
  isAsync: PropTypes.bool.isRequired,
  hasSucceeded: (props, propName) => {
    if (props['isAsync'] && typeof props[propName] !== 'boolean') {
      return new Error('Async WizardPages must have hasSucceeded boolean prop');
    }
  },
  error: PropTypes.object,
  pageList: PropTypes.arrayOf(PropTypes.string).isRequired,
  pageKey: PropTypes.string.isRequired,
  pageIsValid: PropTypes.bool,
  pageIsDirty: PropTypes.bool,
  push: PropTypes.func,
  match: PropTypes.object, //from withRouter
  additionalParams: PropTypes.object,
  windowWidth: PropTypes.number,
};

WizardPage.defaultProps = {
  isAsync: false,
  pageIsValid: true,
  pageIsDirty: true,
};

function mapDispatchToProps(dispatch) {
  return bindActionCreators({ push }, dispatch);
}

const wizardFormPageWithSize = windowSize(WizardPage);
export default withRouter(
  connect(null, mapDispatchToProps)(wizardFormPageWithSize),
);
