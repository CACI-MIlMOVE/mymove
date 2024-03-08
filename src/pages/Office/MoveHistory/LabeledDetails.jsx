/* eslint-disable camelcase */
import React from 'react';

import styles from './LabeledDetails.module.scss';

import booleanFields from 'constants/MoveHistory/Database/BooleanFields';
import dateFields from 'constants/MoveHistory/Database/DateFields';
import fieldMappings from 'constants/MoveHistory/Database/FieldMappings';
import distanceFields from 'constants/MoveHistory/Database/DistanceFields';
import timeUnitFields from 'constants/MoveHistory/Database/TimeUnitFields';
import weightFields from 'constants/MoveHistory/Database/WeightFields';
import monetaryFields from 'constants/MoveHistory/Database/MonetaryFields';
import { shipmentTypes } from 'constants/shipments';
import { HistoryLogRecordShape } from 'constants/MoveHistory/UIDisplay/HistoryLogShape';
import optionFields from 'constants/MoveHistory/Database/OptionFields';
import {
  formatCents,
  formatCustomerDate,
  formatDistanceUnitMiles,
  formatTimeUnitDays,
  formatWeight,
  formatYesNoMoveHistoryValue,
  toDollarString,
} from 'utils/formatters';

export const withMappings = () => {
  const self = {
    displayMappings: [],
  };

  const getResult = ([mapping, fn] = []) => ({
    mapping,
    fn,
  });

  const defaultField = [optionFields, ({ value }) => optionFields[value] || `${value}` || '—'];

  self.addNameMappings = (mappings) => {
    self.displayMappings = self.displayMappings.concat(mappings);
    return self;
  };
  self.getMappedDisplayName = (field) =>
    getResult(self.displayMappings.find(([mapping]) => field in mapping) || defaultField);
  return self;
};

export const { displayMappings, getMappedDisplayName } = withMappings().addNameMappings([
  [weightFields, ({ value }) => formatWeight(Number(value))],
  [dateFields, ({ value }) => formatCustomerDate(value)],
  [booleanFields, ({ value }) => formatYesNoMoveHistoryValue(value)],
  [monetaryFields, ({ value }) => toDollarString(formatCents(value))],
  [timeUnitFields, ({ value }) => formatTimeUnitDays(value)],
  [distanceFields, ({ value }) => formatDistanceUnitMiles(value)],
  [optionFields, ({ value }) => optionFields[value]],
]);

export const retrieveTextToDisplay = (fieldName, value) => {
  const displayName = fieldMappings[fieldName];

  const { fn: valueFormatFn } = getMappedDisplayName(fieldName);
  const displayValue = valueFormatFn({ value });

  return {
    displayName,
    displayValue,
  };
};

// testable for code coverage //
export const createLineItemLabel = (shipmentType, shipmentIdDisplay, serviceItemName, movingExpenseType) =>
  [shipmentType && `${shipmentTypes[shipmentType]} shipment #${shipmentIdDisplay}`, serviceItemName, movingExpenseType]
    .filter((e) => e)
    .join(', ');

// testable for code coverage //

// Filter out empty values unless they used to be non-empty
// These values may be non-nullish in oldValues and nullish in changedValues
// Use the existing keys in changed or old values to check against keys listed in fieldMappings
export const filterInLineItemValues = (changed, old) =>
  Object.entries({ ...old, ...changed }).filter(
    ([theField, theValue]) =>
      !!(
        fieldMappings[theField] &&
        // if changeValues has theField while theValue is either blank or theField exists the old values
        ((theField in changed && (`${theValue}` || theField in old)) ||
          // or theField exists in changeValues
          // and oldValues is empty or not empty
          (theField in changed && (!`${old[theField]}` || `${theValue}`)))
      ),
  );

const LabeledDetails = ({ historyRecord }) => {
  const { changedValues, oldValues = {} } = historyRecord;

  const { shipment_type, shipment_id_display, service_item_name, moving_expense_type, ...changedValuesToUse } =
    changedValues;

  // Check for shipment_type to use it as a header for the row
  const shipmentDisplay = createLineItemLabel(
    shipment_type,
    shipment_id_display,
    service_item_name,
    moving_expense_type,
  );

  const lineItems = filterInLineItemValues(changedValuesToUse, oldValues).map(([label, value]) => {
    const { displayName, displayValue } = retrieveTextToDisplay(label, value);

    return (
      <div key={label}>
        <b>{displayName}</b>: {displayValue}
      </div>
    );
  });

  return (
    <>
      <span className={styles.shipmentType}>{shipmentDisplay}</span>
      {lineItems}
    </>
  );
};

LabeledDetails.propTypes = {
  historyRecord: HistoryLogRecordShape,
};

LabeledDetails.defaultProps = {
  historyRecord: {},
};

export default LabeledDetails;
