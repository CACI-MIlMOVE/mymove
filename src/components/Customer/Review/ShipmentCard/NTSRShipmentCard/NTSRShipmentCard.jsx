import React from 'react';
import { string, shape, func, bool } from 'prop-types';
import { Button } from '@trussworks/react-uswds';
import { generatePath } from 'react-router-dom';

import styles from '../ShipmentCard.module.scss';
import DeliveryDisplay from '../DeliveryDisplay';

import { AddressShape } from 'types/address';
import { getShipmentTypeLabel } from 'utils/shipmentDisplay';
import ShipmentContainer from 'components/Office/ShipmentContainer/ShipmentContainer';
import { customerRoutes } from 'constants/routes';

const NTSRShipmentCard = ({
  destinationLocation,
  destinationZIP,
  secondaryDeliveryAddress,
  receivingAgent,
  remarks,
  requestedDeliveryDate,
  moveId,
  onEditClick,
  onDeleteClick,
  shipmentId,
  shipmentType,
  showEditAndDeleteBtn,
}) => {
  const editPath = generatePath(customerRoutes.SHIPMENT_EDIT_PATH, {
    moveId,
    mtoShipmentId: shipmentId,
  });

  return (
    <div className={styles.ShipmentCard} data-testid="ntsr-summary">
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
              <Button data-testid="edit-ntsr-shipment-btn" onClick={() => onEditClick(editPath)} unstyled>
                Edit
              </Button>
            </div>
          )}
        </div>
        <dl className={styles.shipmentCardSubsection}>
          <DeliveryDisplay
            shipmentId={shipmentId}
            shipmentType={shipmentType}
            requestedDeliveryDate={requestedDeliveryDate}
            destinationLocation={destinationLocation}
            destinationZIP={destinationZIP}
            secondaryDeliveryAddress={secondaryDeliveryAddress}
            receivingAgent={receivingAgent}
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

NTSRShipmentCard.propTypes = {
  destinationLocation: AddressShape,
  destinationZIP: string.isRequired,
  secondaryDeliveryAddress: AddressShape,
  moveId: string.isRequired,
  onEditClick: func.isRequired,
  onDeleteClick: func.isRequired,
  requestedDeliveryDate: string.isRequired,
  showEditAndDeleteBtn: bool.isRequired,
  shipmentId: string.isRequired,
  shipmentType: string.isRequired,
  receivingAgent: shape({
    firstName: string,
    lastName: string,
    phone: string,
    email: string,
  }),
  remarks: string,
};

NTSRShipmentCard.defaultProps = {
  destinationLocation: null,
  receivingAgent: null,
  remarks: '',
  secondaryDeliveryAddress: null,
};

export default NTSRShipmentCard;
