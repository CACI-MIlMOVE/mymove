import React from 'react';
import { string, shape, number, func, bool } from 'prop-types';
import { Button } from '@trussworks/react-uswds';
import { generatePath } from 'react-router-dom';

import styles from '../ShipmentCard.module.scss';
import PickupDisplay from '../PickupDisplay';

import { AddressShape } from 'types/address';
import ShipmentContainer from 'components/Office/ShipmentContainer/ShipmentContainer';
import { getShipmentTypeLabel } from 'utils/shipmentDisplay';
import { customerRoutes } from 'constants/routes';

const NTSShipmentCard = ({
  moveId,
  onEditClick,
  onDeleteClick,
  pickupLocation,
  secondaryPickupAddress,
  releasingAgent,
  remarks,
  requestedPickupDate,
  shipmentId,
  shipmentType,
  shipmentNumber,
  showEditAndDeleteBtn,
}) => {
  const editPath = generatePath(customerRoutes.SHIPMENT_EDIT_PATH, {
    moveId,
    mtoShipmentId: shipmentId,
  });

  return (
    <div className={styles.ShipmentCard} data-testid="nts-summary">
      <ShipmentContainer className={styles.container} shipmentType={shipmentType}>
        <div className={styles.ShipmentCardHeader}>
          <div className={styles.shipmentTypeNumber}>
            <h3>{getShipmentTypeLabel(shipmentType)}</h3>
            <p>#{shipmentId.substring(0, 8).toUpperCase()}</p>
          </div>
          {showEditAndDeleteBtn && (
            <div className={styles.btnContainer}>
              <Button onClick={() => onDeleteClick(shipmentId)} unstyled>
                Delete
              </Button>
              |
              <Button data-testid="edit-nts-shipment-btn" onClick={() => onEditClick(editPath)} unstyled>
                Edit
              </Button>
            </div>
          )}
        </div>
        <dl className={styles.shipmentCardSubsection}>
          <PickupDisplay
            shipmentId={shipmentId}
            shipmentType={shipmentType}
            shipmentNumber={shipmentNumber}
            requestedPickupDate={requestedPickupDate}
            pickupLocation={pickupLocation}
            secondaryPickupAddress={secondaryPickupAddress}
            releasingAgent={releasingAgent}
          />
          <div className={`${styles.row} ${styles.remarksRow}`}>
            <dt>Remarks</dt>
            <dd className={styles.remarksCell}>{remarks || '—'}</dd>
          </div>
        </dl>
      </ShipmentContainer>
    </div>
  );
};

NTSShipmentCard.propTypes = {
  moveId: string.isRequired,
  onEditClick: func.isRequired,
  onDeleteClick: func.isRequired,
  shipmentType: string.isRequired,
  shipmentId: string.isRequired,
  showEditAndDeleteBtn: bool.isRequired,
  requestedPickupDate: string.isRequired,
  pickupLocation: AddressShape.isRequired,
  secondaryPickupAddress: AddressShape,
  releasingAgent: shape({
    firstName: string,
    lastName: string,
    phone: string,
    email: string,
  }),
  remarks: string,
  shipmentNumber: number,
};

NTSShipmentCard.defaultProps = {
  releasingAgent: null,
  remarks: '',
  shipmentNumber: 0,
  secondaryPickupAddress: null,
};

export default NTSShipmentCard;
