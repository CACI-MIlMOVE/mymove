import React from 'react';
import classNames from 'classnames/bind';
import PropTypes from 'prop-types';

import styles from './SectionWrapper.module.scss';

const cx = classNames.bind(styles);

const SectionWrapper = ({ children }) => <div className={cx('sectionWrapper')}>{children}</div>;

SectionWrapper.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.element), PropTypes.element.isRequired]),
};

SectionWrapper.defaultProps = {
  children: ' ',
};

export { SectionWrapper };
