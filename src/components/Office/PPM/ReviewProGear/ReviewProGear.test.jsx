import React from 'react';
import { render, waitFor, screen, fireEvent } from '@testing-library/react';

import ReviewProGear from './ReviewProGear';

import ppmDocumentStatus from 'constants/ppms';
import { MockProviders } from 'testUtils';

beforeEach(() => {
  jest.clearAllMocks();
});

const defaultProps = {
  ppmShipment: {
    id: '32ecb311-edbe-4fd4-96ee-bd693113f3f3',
    actualMoveDate: '02-Dec-22',
    actualPickupPostalCode: '90210',
    actualDestinationPostalCode: '94611',
    hasReceivedAdvance: true,
    advanceAmountReceived: 60000,
  },
  tripNumber: 1,
  ppmNumber: 1,
};

const proGearRequiredProps = {
  proGear: {
    id: '32ecb311-edbe-4fd4-96ee-bd693113f3f3',
    belongsToSelf: true,
    weight: 400,
    description: 'Kia Forte',
    hasWeightTickets: true,
  },
};

const missingWeightTicketProps = {
  proGear: {
    id: '32ecb311-edbe-4fd4-96ee-bd693113f3f3',
    belongsToSelf: true,
    weight: 400,
    description: 'Kia Forte',
    hasWeightTickets: false,
  },
};

const rejectedProps = {
  proGear: {
    ...proGearRequiredProps.proGear,
    status: ppmDocumentStatus.REJECTED,
    reason: 'Rejection reason',
  },
};

describe('ReviewProGear component', () => {
  describe('displays form', () => {
    it('renders blank form on load with defaults', async () => {
      render(
        <MockProviders>
          <ReviewProGear {...defaultProps} />;
        </MockProviders>,
      );

      await waitFor(() => {
        expect(screen.getByRole('heading', { level: 3, name: 'Pro-gear 1' })).toBeInTheDocument();
      });
      expect(screen.getByText('Belongs to')).toBeInTheDocument();
      expect(screen.getByLabelText('Customer')).toBeInstanceOf(HTMLInputElement);
      expect(screen.getByLabelText('Spouse')).toBeInstanceOf(HTMLInputElement);

      expect(screen.getByText('Description')).toBeInTheDocument();

      expect(screen.getByText('Pro-gear weight')).toBeInTheDocument();
      expect(screen.getByLabelText('Weight tickets')).toBeInstanceOf(HTMLInputElement);
      expect(screen.getByLabelText('Constructed weight')).toBeInstanceOf(HTMLInputElement);

      expect(screen.getByLabelText('Constructed pro-gear weight')).toBeInstanceOf(HTMLInputElement);

      expect(screen.getByRole('heading', { level: 3, name: 'Review pro-gear 1' })).toBeInTheDocument();
      expect(screen.getByText('Add a review for this pro-gear')).toBeInTheDocument();
      expect(screen.getByLabelText('Accept')).toBeInstanceOf(HTMLInputElement);
      expect(screen.getByLabelText('Reject')).toBeInstanceOf(HTMLInputElement);
    });

    it('populates edit form with existing pro-gear weight ticket values', async () => {
      render(
        <MockProviders>
          <ReviewProGear {...defaultProps} {...proGearRequiredProps} />;
        </MockProviders>,
      );

      await waitFor(() => {
        expect(screen.getByLabelText('Customer')).toBeChecked();
      });
      expect(screen.getByText('Kia Forte')).toBeInTheDocument();
      expect(screen.getByLabelText(/Shipment's pro-gear weight/)).toHaveDisplayValue('400');
    });

    it('populates edit form when pro-gear weight ticket is missing', async () => {
      render(
        <MockProviders>
          <ReviewProGear {...defaultProps} {...missingWeightTicketProps} />;
        </MockProviders>,
      );
      await waitFor(() => {
        expect(screen.getByLabelText('Constructed weight')).toBeChecked();
        expect(screen.getByText('Constructed pro-gear weight')).toBeInTheDocument();
      });
    });

    it('displays remaining character count', async () => {
      render(
        <MockProviders>
          <ReviewProGear {...defaultProps} {...rejectedProps} />
        </MockProviders>,
      );
      await waitFor(() => {
        expect(screen.getByLabelText('Reason')).toHaveDisplayValue('Rejection reason');
      });
      expect(screen.getByText('484 characters')).toBeInTheDocument();
    });

    it('toggles the reason field when Reject is selected', async () => {
      render(
        <MockProviders>
          <ReviewProGear {...defaultProps} {...proGearRequiredProps} />
        </MockProviders>,
      );
      await waitFor(() => {
        expect(screen.getByLabelText('Reject')).toBeInstanceOf(HTMLInputElement);
      });
      await fireEvent.click(screen.getByLabelText('Reject'));
      expect(screen.getByLabelText('Reason')).toBeInstanceOf(HTMLTextAreaElement);
      await fireEvent.click(screen.getByLabelText('Accept'));
      expect(screen.queryByLabelText('Reason')).not.toBeInTheDocument();
    });
  });
});
