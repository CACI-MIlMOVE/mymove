import React from 'react';
import { action } from '@storybook/addon-actions';
import { Grid, GridContainer } from '@trussworks/react-uswds';
import { within, userEvent } from '@storybook/testing-library';

import EstimatedWeightsProGear from './EstimatedWeightsProGear';

export default {
  title: 'Customer Components / PPM Booking / Estimated Weights and Pro-gear',
  component: EstimatedWeightsProGear,
  decorators: [
    (Story) => (
      <GridContainer>
        <Grid row>
          <Grid col desktop={{ col: 8, offset: 2 }}>
            <Story />
          </Grid>
        </Grid>
      </GridContainer>
    ),
  ],
};

const Template = (args) => <EstimatedWeightsProGear {...args} />;

export const BlankEstimatedWeightsProGear = Template.bind({});
BlankEstimatedWeightsProGear.args = {
  onSubmit: action('submit button clicked'),
  onBack: action('back button clicked'),
  entitlement: {
    authorizedWeight: 5000,
  },
};

export const MTOShipmentEstimatedWeightProGear = Template.bind({});
MTOShipmentEstimatedWeightProGear.args = {
  onSubmit: action('submit button clicked'),
  onBack: action('back button clicked'),
  entitlement: {
    authorizedWeight: 5000,
  },
  mtoShipment: {
    id: '123',
    ppmShipment: {
      id: '123',
      hasProGear: 'true',
      estimatedProGearWeight: '1000',
      estimatedSpouseProGearWeight: '100',
      estimatedWeight: '4000',
    },
  },
};

export const ErrorEstimatedWeightsProGear = Template.bind({});
ErrorEstimatedWeightsProGear.args = {
  onSubmit: action('submit button clicked'),
  onBack: action('back button clicked'),
  entitlement: {
    authorizedWeight: 5000,
  },
};
ErrorEstimatedWeightsProGear.play = async ({ canvasElement }) => {
  // Starts querying the component from its root element
  const canvas = within(canvasElement);

  // See https://storybook.js.org/docs/react/essentials/actions#automatically-matching-args to learn how to setup logging in the Actions panel
  await userEvent.click(canvas.getByText('Save & Continue'));
};
