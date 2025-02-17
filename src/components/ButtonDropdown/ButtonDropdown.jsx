import React from 'react';
import { Dropdown } from '@trussworks/react-uswds';
import PropTypes from 'prop-types';
import classnames from 'classnames';

import styles from './ButtonDropdown.module.scss';

const ButtonDropdown = ({ children, onChange, ariaLabel, divClassName }) => (
  <div className={classnames(styles.ButtonDropdown, divClassName)}>
    <Dropdown aria-label={ariaLabel} onChange={onChange} className={classnames(styles.ButtonDropdown, 'usa-button')}>
      {children}
    </Dropdown>
    <span className={styles.ButtonDropdownIcon} />
  </div>
);

ButtonDropdown.defaultProps = {
  ariaLabel: '',
  divClassName: '',
};

ButtonDropdown.propTypes = {
  children: PropTypes.node.isRequired,
  onChange: PropTypes.func.isRequired,
  ariaLabel: PropTypes.string,
  divClassName: PropTypes.string,
};

export default ButtonDropdown;
