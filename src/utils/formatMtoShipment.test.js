import {
  formatMtoShipmentForAPI,
  formatMtoShipmentForDisplay,
  formatPpmShipmentForAPI,
  formatPpmShipmentForDisplay,
  getMtoShipmentLabel,
} from './formatMtoShipment';

import { MTOAgentType, SHIPMENT_OPTIONS } from 'shared/constants';

describe('formatMtoShipmentForDisplay', () => {
  const emptyAgentShape = {
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
  };

  const emptyAddressShape = {
    streetAddress1: '',
    streetAddress2: '',
    city: '',
    state: '',
    postalCode: '',
  };

  const mtoShipment = {
    requestedPickupDate: '2026-01-20',
    pickupAddress: {
      streetAddress1: '123 main',
      city: 'legit human city',
      state: 'DC',
      postalCode: '20017',
    },
    requestedDeliveryDate: '2026-01-27',
    moveTaskOrderID: 'move123',
  };

  const customerRemarks = 'some mock remarks';
  const counselorRemarks = 'all looks good';

  const releasingAgent = {
    firstName: 'mockFirstName',
    lastName: 'mockLastName',
    email: 'mockAgentEmail@example.com',
    phone: '222-555-1234',
  };

  const receivingAgent = {
    firstName: 'r0b0tBestFr1end',
    lastName: 'r0b0tBestFr1endLastName',
    email: 'r0b0t-fr1end@example.com',
    phone: '222-555-0101',
  };

  const destinationAddress = {
    streetAddress1: '0011100010110101',
    city: 'R0B0T T0WN',
    state: 'CP',
    postalCode: '10101',
  };

  const secondaryPickupAddress = {
    streetAddress1: '142 E Barrel Hoop Circle',
    streetAddress2: '#4A',
    city: 'Corpus Christi',
    state: 'TX',
    postalCode: '78412',
  };

  const secondaryDeliveryAddress = {
    streetAddress1: '441 SW Río de la Plata Drive',
    streetAddress2: '',
    city: destinationAddress.city,
    state: destinationAddress.state,
    postalCode: destinationAddress.postalCode,
  };

  const checkAddressesAreEqual = (address1, address2) => {
    expect(address1.streetAddress1 === address2.streetAddress1);
    expect(address1.streetAddress2 === address2.streetAddress2);
    expect(address1.streetAddress3 === address2.streetAddress3);
    expect(address1.city === address2.city);
    expect(address1.state === address2.state);
    expect(address1.postalCode === address2.postalCode);
  };

  const checkAgentsAreEqual = (agent1, agent2) => {
    expect(agent1.firstName === agent2.firstName);
    expect(agent1.lastName === agent2.lastName);
    expect(agent1.email === agent2.email);
    expect(agent1.phone === agent2.phone);
    expect(agent1.agentType === agent2.agentType);
  };

  it.each([[SHIPMENT_OPTIONS.HHG], [SHIPMENT_OPTIONS.NTSR], [SHIPMENT_OPTIONS.NTS]])(
    'can format a shipment (type: %s)',
    (shipmentType) => {
      const params = {
        ...mtoShipment,
        shipmentType,
      };

      const displayValues = formatMtoShipmentForDisplay(params);

      expect(displayValues.shipmentType).toBe(shipmentType);
      expect(displayValues.moveTaskOrderID).toBe(mtoShipment.moveTaskOrderID);
      expect(displayValues.customerRemarks).toBe('');
      expect(displayValues.counselorRemarks).toBe('');

      expect(displayValues.pickup.requestedDate.toDateString()).toBe('Tue Jan 20 2026');

      const expectedPickupAddress = { ...emptyAddressShape, ...mtoShipment.pickupAddress };
      checkAddressesAreEqual(displayValues.pickup.address, expectedPickupAddress);
      checkAgentsAreEqual(displayValues.pickup.agent, emptyAgentShape);

      expect(displayValues.delivery.requestedDate.toDateString()).toBe('Tue Jan 27 2026');
      checkAddressesAreEqual(displayValues.delivery.address, emptyAddressShape);
      checkAgentsAreEqual(displayValues.delivery.agent, emptyAgentShape);
      expect(displayValues.hasDeliveryAddress).toBe('no');

      checkAddressesAreEqual(displayValues.secondaryPickup.address, emptyAddressShape);
      expect(displayValues.hasSecondaryPickup).toBe('no');

      checkAddressesAreEqual(displayValues.secondaryDelivery.address, emptyAddressShape);
      expect(displayValues.hasSecondaryDelivery).toBe('no');

      expect(displayValues.agents).toBeUndefined();
    },
  );

  it('can format a shipment with remarks', () => {
    const params = {
      ...mtoShipment,
      shipmentType: SHIPMENT_OPTIONS.HHG,
      customerRemarks,
      counselorRemarks,
    };

    const displayValues = formatMtoShipmentForDisplay(params);

    expect(displayValues.customerRemarks).toBe(customerRemarks);
    expect(displayValues.counselorRemarks).toBe(counselorRemarks);
  });

  it('can format a shipment with agents', () => {
    const params = {
      ...mtoShipment,
      shipmentType: SHIPMENT_OPTIONS.HHG,
      agents: [
        { ...releasingAgent, agentType: MTOAgentType.RELEASING },
        { ...receivingAgent, agentType: MTOAgentType.RECEIVING },
      ],
    };

    const displayValues = formatMtoShipmentForDisplay(params);

    checkAgentsAreEqual(displayValues.pickup.agent, releasingAgent);
    checkAgentsAreEqual(displayValues.delivery.agent, receivingAgent);
  });

  it('can format a shipment with a destination, secondary pickup, and secondary destination', () => {
    const params = {
      ...mtoShipment,
      shipmentType: SHIPMENT_OPTIONS.HHG,
      destinationAddress,
      secondaryPickupAddress,
      secondaryDeliveryAddress,
    };

    const displayValues = formatMtoShipmentForDisplay(params);

    const expectedDeliveryAddress = { ...emptyAddressShape, ...destinationAddress };
    checkAddressesAreEqual(displayValues.delivery.address, expectedDeliveryAddress);
    expect(displayValues.hasDeliveryAddress).toBe('yes');

    const expectedSecondaryPickupAddress = { ...emptyAddressShape, ...secondaryPickupAddress };
    checkAddressesAreEqual(displayValues.secondaryPickup.address, expectedSecondaryPickupAddress);
    expect(displayValues.hasSecondaryPickup).toBe('yes');

    const expectedSecondaryDeliveryAddress = { ...emptyAddressShape, ...secondaryDeliveryAddress };
    checkAddressesAreEqual(displayValues.secondaryDelivery.address, expectedSecondaryDeliveryAddress);
    expect(displayValues.hasSecondaryDelivery).toBe('yes');
  });

  it('can format a shipment with lines of accounting', () => {
    const params = {
      ...mtoShipment,
      tacType: 'HHG',
      sacType: 'NTS',
    };

    const displayValues = formatMtoShipmentForDisplay(params);
    expect(displayValues.tacType).toEqual('HHG');
    expect(displayValues.sacType).toEqual('NTS');
  });

  it('can format a shipment with shipment weight', () => {
    const params = {
      ...mtoShipment,
      ntsRecordedWeight: 4000,
    };

    const displayValues = formatMtoShipmentForDisplay(params);
    expect(displayValues.ntsRecordedWeight).toEqual(4000);
  });
});

describe('formatMtoShipmentForAPI', () => {
  const mtoShipmentParams = {
    moveId: 'move123',
    customerRemarks: 'some mock remarks',
  };

  const pickupInfo = {
    requestedDate: '2026-01-07',
    address: {
      streetAddress1: '123 main',
      city: 'legit human city',
      state: 'DC',
      postalCode: '20017',
    },
    agent: {
      firstName: 'mockFirstName',
      lastName: 'mockLastName',
      email: 'mockAgentEmail@example.com',
      phone: '222-555-1234',
    },
  };

  const deliveryInfo = {
    requestedDate: '2026-01-27',
    address: {
      streetAddress1: '0011100010110101',
      city: 'R0B0T T0WN',
      state: 'CP',
      postalCode: '10101',
    },
    agent: {
      firstName: 'r0b0tBestFr1end',
      lastName: 'r0b0tBestFr1endLastName',
      email: 'r0b0t-fr1end@example.com',
      phone: '222-555-0101',
    },
  };

  const storageFacility = {
    facilityName: 'Most Excellent Storage',
    phone: '999-999-9999',
    lotNumber: 42,
    address: {
      streetAddress1: '3373 NW Martin Luther King Blvd',
      city: 'San Antonio',
      state: 'TX',
      ZIP: '78234',
      eTag: '678',
    },
    eTag: '456',
  };

  it('can format an HHG shipment', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.HHG,
      pickup: { ...pickupInfo },
      delivery: { ...deliveryInfo },
    };

    const actual = formatMtoShipmentForAPI(params);

    expect(actual.shipmentType).toBe(SHIPMENT_OPTIONS.HHG);
    expect(actual.agents.length).toBe(2);
    expect(actual.requestedPickupDate).toBe('2026-01-07');
    expect(actual.agents[0].phone).toBe('222-555-1234');
    expect(actual.agents[0].agentType).toBe('RELEASING_AGENT');
    expect(actual.requestedDeliveryDate).toBe('2026-01-27');
    expect(actual.agents[1].phone).toBe('222-555-0101');
    expect(actual.agents[1].agentType).toBe('RECEIVING_AGENT');
    expect(actual.customerRemarks).toBe('some mock remarks');

    expect(actual.secondaryPickupAddress).toBeUndefined();
    expect(actual.secondaryDeliveryAddress).toBeUndefined();
  });

  it('can format an NTSr shipment', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.NTSR,
      pickup: { ...pickupInfo },
      storageFacility,
    };

    const actual = formatMtoShipmentForAPI(params);

    expect(actual.shipmentType).toBe(SHIPMENT_OPTIONS.NTSR);
    expect(actual.requestedPickupDate).toBe('2026-01-07');
    expect(actual.agents.length).toBe(1);
    expect(actual.agents[0].phone).toBe('222-555-1234');
    expect(actual.agents[0].agentType).toBe('RELEASING_AGENT');
    expect(actual.customerRemarks).toBe('some mock remarks');

    expect(actual.storageFacility.eTag).toBeUndefined();
    expect(actual.storageFacility.address.eTag).toBeUndefined();
  });

  it('can format an NTS shipment', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.NTS,
      delivery: { ...deliveryInfo },
    };
    const actual = formatMtoShipmentForAPI(params);
    expect(actual.shipmentType).toBe(SHIPMENT_OPTIONS.NTS);
    expect(actual.requestedDeliveryDate).toBe('2026-01-27');
    expect(actual.agents.length).toBe(1);
    expect(actual.agents[0].phone).toBe('222-555-0101');
    expect(actual.agents[0].agentType).toBe('RECEIVING_AGENT');
    expect(actual.customerRemarks).toBe('some mock remarks');
  });

  it('can format a shipment with lines of accounting', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.NTS,
      pickup: { ...pickupInfo },
      tacType: 'HHG',
      sacType: 'NTS',
    };

    const actual = formatMtoShipmentForAPI(params);

    expect(actual.tacType).toEqual('HHG');
    expect(actual.sacType).toEqual('NTS');
  });

  it('can format a shipment with shipment weight', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.NTSR,
      delivery: { ...deliveryInfo },
      ntsRecordedWeight: '4000',
      storageFacility,
    };

    const actual = formatMtoShipmentForAPI(params);
    expect(actual.ntsRecordedWeight).toEqual(4000);
  });

  it('can format a shipment with shipment weight including delimiters', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.NTSR,
      delivery: { ...deliveryInfo },
      ntsRecordedWeight: '4,500',
      storageFacility,
    };

    const actual = formatMtoShipmentForAPI(params);
    expect(actual.ntsRecordedWeight).toEqual(4500);
  });

  it('can format an HHG shipment with a secondary pickup/destination', () => {
    const params = {
      ...mtoShipmentParams,
      shipmentType: SHIPMENT_OPTIONS.HHG,
      pickup: { ...pickupInfo },
      hasSecondaryPickup: true,
      secondaryPickup: {
        address: {
          streetAddress1: '142 E Barrel Hoop Circle',
          streetAddress2: '#4A',
          city: 'Corpus Christi',
          state: 'TX',
          postalCode: '78412',
        },
      },
      delivery: { ...deliveryInfo },
      hasSecondaryDelivery: true,
      secondaryDelivery: {
        address: {
          streetAddress1: '441 SW Río de la Plata Drive',
          streetAddress2: '',
          city: deliveryInfo.address.city,
          state: deliveryInfo.address.state,
          postalCode: deliveryInfo.address.postalCode,
        },
      },
    };

    const actual = formatMtoShipmentForAPI(params);

    expect(actual.secondaryPickupAddress).not.toBeUndefined();
    expect(actual.secondaryPickupAddress.streetAddress1).toEqual('142 E Barrel Hoop Circle');

    expect(actual.secondaryDeliveryAddress).not.toBeUndefined();
    expect(actual.secondaryDeliveryAddress.streetAddress1).toEqual('441 SW Río de la Plata Drive');
  });
});

describe('formatPpmShipmentForDisplay', () => {
  it('creates a base display values object without an existing shipment', () => {
    const display = formatPpmShipmentForDisplay({});

    expect(display.estimatedWeight).toEqual('');
    expect(display.hasProGear).toEqual(false);
    expect(display.advanceRequested).toEqual(false);
  });

  it('converts an existing shipment to formatted display values', () => {
    const api = {
      expectedDepatureDate: '2022-12-25',
      pickupPostalCode: '90210',
      secondaryPickupPostalCode: '80014',
      destinationPostalCode: '22201',
      secondaryDestinationPostalCode: '20002',

      sitExpected: true,
      sitLocation: 'DESTINATION',
      sitEstimatedWeight: 2750,
      sitEstimatedEntryDate: '2022-12-01',
      sitEstimatedDepartureDate: '2022-12-15',

      estimatedWeight: 9000,
      hasProGear: true,
      proGearWeight: 1000,

      estimatedIncentive: 400000,
      hasRequestedAdvance: true,
      advanceAmountRequested: 200000,
    };

    const display = formatPpmShipmentForDisplay({ ppmShipment: api, counselorRemarks: 'test remarks' });

    expect(display.secondPickupPostalCode).toEqual('80014');
    expect(display.sitEstimatedWeight).toEqual('2750');
    expect(display.estimatedWeight).toEqual('9000');
    expect(display.proGearWeight).toEqual('1000');
    expect(display.advanceRequested).toEqual(true);
    expect(display.advance).toEqual('2000');
    expect(display.counselorRemarks).toEqual('test remarks');
  });
});

describe('formatPpmShipmentForAPI', () => {
  it('converts fully filled-out formValues to api values', () => {
    const formValues = {
      expectedDepatureDate: '2022-12-25',
      pickupPostalCode: '90210',
      secondPickupPostalCode: '80014',
      destinationPostalCode: '22201',

      sitExpected: true,
      sitLocation: 'ORIGIN',
      sitEstimatedWeight: '2500',
      sitEstimatedEntryDate: '2022-12-01',
      sitEstimatedDepartureDate: '2022-12-15',

      estimatedWeight: '7500',
      hasProGear: true,
      proGearWeight: '1000',

      advanceRequested: true,
      advance: '2000',

      counselorRemarks: 'test remarks',
    };

    const { counselorRemarks, ppmShipment } = formatPpmShipmentForAPI(formValues);

    expect(ppmShipment.secondaryPickupPostalCode).toEqual('80014');
    expect(ppmShipment.estimatedWeight).toEqual(7500);
    expect(ppmShipment.proGearWeight).toEqual(1000);
    expect(ppmShipment.spouseProGearWeight).toEqual(undefined);

    expect(ppmShipment.hasRequestedAdvance).toEqual(true);
    expect(ppmShipment.advanceAmountRequested).toEqual(200000);

    expect(counselorRemarks).toEqual('test remarks');
  });

  it('converts minimal formValues to api values', () => {
    const formValues = {
      expectedDepatureDate: '2022-12-25',
      pickupPostalCode: '90210',
      destinationPostalCode: '22201',

      sitExpected: false,

      estimatedWeight: '7500',
      hasProGear: false,
    };

    const { ppmShipment } = formatPpmShipmentForAPI(formValues);

    expect(ppmShipment.estimatedWeight).toEqual(7500);

    expect(ppmShipment.sitLocation).toEqual(undefined);
    expect(ppmShipment.proGearWeight).toEqual(undefined);
  });
});
describe('getMtoShipmentLabel', () => {
  const historyRecord = {
    changedValues: {
      status: 'SUBMITTED',
    },
  };
  const context = [
    {
      shipment_type: 'HHG',
      shipment_id_abbr: 'a1a1a',
      name: 'Bars',
    },
  ];
  const contextNoShipmentType = [
    {
      shipment_id_abbr: 'a1a1a',
      name: 'Bars',
    },
  ];
  const contextNoShipmentId = [
    {
      shipment_type: 'HHG',
      name: 'Bars',
    },
  ];
  const contextNoServiceItem = [
    {
      shipment_type: 'HHG',
      shipment_id_abbr: 'a1a1a',
    },
  ];
  it('returns an empty object if context is not present', () => {
    const result = getMtoShipmentLabel(historyRecord);
    expect(result).toEqual({});
  });
  it('returns information need to generate shipment label used in move history', () => {
    const result = getMtoShipmentLabel({ ...historyRecord, context });
    expect(result).toEqual({ shipment_type: 'HHG', shipment_id_display: 'A1A1A', service_item_name: 'Bars' });
  });
  it('returns object without shipment_type when shipment_type is not present in context', () => {
    const result = getMtoShipmentLabel({ ...historyRecord, context: contextNoShipmentType });
    expect(result).toEqual({ shipment_id_display: 'A1A1A', service_item_name: 'Bars' });
  });
  it('returns object without shipment_id_display when shipment_id_abbr is not present in context', () => {
    const result = getMtoShipmentLabel({ ...historyRecord, context: contextNoShipmentId });
    expect(result).toEqual({ shipment_type: 'HHG', service_item_name: 'Bars' });
  });
  it('returns object without shipment_service_item_name when name is not present in context ', () => {
    const result = getMtoShipmentLabel({ ...historyRecord, context: contextNoServiceItem });
    expect(result).toEqual({ shipment_type: 'HHG', shipment_id_display: 'A1A1A' });
  });
});
