import { selectEntitlements } from './entitlements';

describe('entitlements', () => {
  describe('when I have dependents', () => {
    describe('when my spouse has pro gear', () => {
      it('should include spouse progear', () => {
        const entitlements = selectEntitlements(
          {
            total_weight_self: 5000,
            total_weight_self_plus_dependents: 8000,
            pro_gear_weight: 2000,
            pro_gear_weight_spouse: 500,
          },
          true,
          true,
        );
        expect(entitlements).toEqual({
          pro_gear: 2000,
          pro_gear_spouse: 500,
          sum: 10500,
          weight: 8000,
        });
      });
    });
    describe('when my spouse does not have pro gear', () => {
      it('should not include spouse progear', () => {
        const entitlements = selectEntitlements(
          {
            total_weight_self: 5000,
            total_weight_self_plus_dependents: 8000,
            pro_gear_weight: 2000,
            pro_gear_weight_spouse: 500,
          },
          true,
          false,
        );
        expect(entitlements).toEqual({
          pro_gear: 2000,
          pro_gear_spouse: 0,
          sum: 10000,
          weight: 8000,
        });
      });
    });
  });
  describe("when I don't have dependents", () => {
    it('should exclude spouse progear', () => {
      const entitlements = selectEntitlements({
        total_weight_self: 5000,
        total_weight_self_plus_dependents: 8000,
        pro_gear_weight: 2000,
        pro_gear_weight_spouse: 500,
      });
      expect(entitlements).toEqual({
        pro_gear: 2000,
        pro_gear_spouse: 0,
        sum: 7000,
        weight: 5000,
      });
    });
  });
});
