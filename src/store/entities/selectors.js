import { createSelector } from 'reselect';

import { profileStates } from 'constants/customerStates';
import { MOVE_STATUSES, NULL_UUID } from 'shared/constants';

/**
 * Use this file for selecting "slices" of state from Redux and for computed
 * properties given state. Selectors can be memoized for performance.
 * Documentation: https://github.com/reduxjs/reselect
 */

/** User */
export const selectLoggedInUser = (state) => {
  if (state.entities.user) return Object.values(state.entities.user)[0];
  return null;
};

/** Service Member */
export const selectServiceMemberFromLoggedInUser = (state) => {
  const user = selectLoggedInUser(state);
  if (!user || !user.service_member) return null;
  return state.entities.serviceMembers?.[`${user.service_member}`] || null;
};

export const selectCurrentDutyLocation = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  return serviceMember?.current_location || null;
};

export const selectServiceMemberAffiliation = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  return serviceMember?.affiliation || null;
};

export const selectServiceMemberProfileState = createSelector(selectServiceMemberFromLoggedInUser, (serviceMember) => {
  if (!serviceMember) return profileStates.EMPTY_PROFILE;

  /* eslint-disable camelcase */
  const {
    rank,
    edipi,
    affiliation,
    first_name,
    last_name,
    telephone,
    personal_email,
    phone_is_preferred,
    email_is_preferred,
    current_location,
    residential_address,
    backup_mailing_address,
    backup_contacts,
  } = serviceMember;

  if (!rank || !edipi || !affiliation) return profileStates.EMPTY_PROFILE;
  if (!first_name || !last_name) return profileStates.DOD_INFO_COMPLETE;
  if (!telephone || !personal_email || !(phone_is_preferred || email_is_preferred)) return profileStates.NAME_COMPLETE;
  if (!current_location || !current_location.id || current_location.id === NULL_UUID)
    return profileStates.CONTACT_INFO_COMPLETE;
  if (!residential_address) return profileStates.DUTY_LOCATION_COMPLETE;
  if (!backup_mailing_address) return profileStates.ADDRESS_COMPLETE;
  if (!backup_contacts || !backup_contacts.length) return profileStates.BACKUP_ADDRESS_COMPLETE;
  return profileStates.BACKUP_CONTACTS_COMPLETE;
  /* eslint-enable camelcase */
});

// TODO: this is similar to service_member.isProfileComplete and we should figure out how to use just one if possible
export const selectIsProfileComplete = createSelector(
  selectServiceMemberFromLoggedInUser,
  (serviceMember) =>
    !!(
      serviceMember &&
      serviceMember.rank &&
      serviceMember.edipi &&
      serviceMember.affiliation &&
      serviceMember.first_name &&
      serviceMember.last_name &&
      serviceMember.telephone &&
      serviceMember.personal_email &&
      serviceMember.current_location?.id &&
      serviceMember.residential_address?.postalCode &&
      serviceMember.backup_mailing_address?.postalCode &&
      serviceMember.backup_contacts?.length > 0
    ),
);

/** Backup Contacts */
export const selectBackupContacts = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  const backupContactIds = serviceMember?.backup_contacts || [];
  return backupContactIds.map((id) => state.entities.backupContacts?.[`${id}`]);
};

/** Orders */
export const selectOrdersForLoggedInUser = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  const ordersIds = serviceMember?.orders || [];
  return ordersIds.map((id) => state.entities.orders?.[`${id}`]);
};

export const selectCurrentOrders = (state) => {
  const orders = selectOrdersForLoggedInUser(state);
  const [activeOrders] = orders.filter(
    (o) => ['DRAFT', 'SUBMITTED', 'APPROVED', 'PAYMENT_REQUESTED'].indexOf(o?.status) > -1,
  );

  return activeOrders || orders[0] || null;
};

export const selectOrdersById = (state, id) => {
  return state.entities.orders?.[`${id}`] || null;
};

export const selectUploadsForCurrentOrders = (state) => {
  const orders = selectCurrentOrders(state);
  return orders ? orders.uploaded_orders?.uploads : [];
};

export const selectUploadsForCurrentAmendedOrders = (state) => {
  const orders = selectCurrentOrders(state);
  return orders ? orders.uploaded_amended_orders?.uploads : [];
};

/** Moves */
export const selectMovesForLoggedInUser = (state) => {
  const orders = selectOrdersForLoggedInUser(state);
  const moves = orders?.reduce((prev, cur) => {
    return prev.concat(cur.moves?.map((id) => state.entities.moves?.[`${id}`]) || []);
  }, []);

  return moves;
};

export const selectMovesForCurrentOrders = (state) => {
  const activeOrders = selectCurrentOrders(state);
  const moveIds = activeOrders?.moves || [];
  return moveIds.map((id) => state.entities.moves?.[`${id}`]);
};

export const selectCurrentMove = (state) => {
  const moves = selectMovesForCurrentOrders(state);
  const [activeMove] = moves.filter(
    (m) => ['DRAFT', 'SUBMITTED', 'APPROVED', 'PAYMENT_REQUESTED'].indexOf(m?.status) > -1,
  );
  return activeMove || moves[0] || null;
};

export const selectMoveIsApproved = createSelector(selectCurrentMove, (move) => move?.status === 'APPROVED');

export const selectMoveIsInDraft = createSelector(selectCurrentMove, (move) => move?.status === MOVE_STATUSES.DRAFT);

export const selectHasCanceledMove = createSelector(selectMovesForLoggedInUser, (moves) =>
  moves.some((m) => m.status === 'CANCELED'),
);

/** MTO Shipments */
export const selectMTOShipmentsForCurrentMove = (state) => {
  const currentMove = selectCurrentMove(state);
  return Object.values(state.entities.mtoShipments)?.filter((m) => m.moveTaskOrderID === currentMove?.id);
};

export function selectMTOShipmentById(state, id) {
  return state.entities?.mtoShipments?.[`${id}`] || null;
}

/** PPMs */
export const selectWeightTicketAndIndexById = (state, mtoShipmentId, weightTicketId) => {
  let weightTicket = null;
  let index = -1;
  if (weightTicketId == null) {
    return { weightTicket, index };
  }

  const mtoShipment = selectMTOShipmentById(state, mtoShipmentId);
  const weightTickets = mtoShipment?.ppmShipment?.weightTickets;
  if (Array.isArray(weightTickets)) {
    index = weightTickets.findIndex((ele) => ele.id === weightTicketId);
    weightTicket = weightTickets?.[index] || null;
  }
  return { weightTicket, index };
};

export const selectWeightTicketsForShipment = (state, mtoShipmentId) => {
  const mtoShipment = selectMTOShipmentById(state, mtoShipmentId);
  return mtoShipment?.ppmShipment?.weightTickets;
};

export const selectExpenseAndIndexById = (state, mtoShipmentId, expenseId) => {
  let expense = null;
  let index = -1;
  if (expenseId == null) {
    return { expense, index };
  }

  const mtoShipment = selectMTOShipmentById(state, mtoShipmentId);
  const expenses = mtoShipment?.ppmShipment?.movingExpenses;
  if (Array.isArray(expenses)) {
    index = expenses.findIndex((el) => el.id === expenseId);
    expense = expenses?.[index] || null;
  }
  return { expense, index };
};

export const selectPPMForMove = (state, moveId) => {
  const ppmForMove = Object.values(state.entities.personallyProcuredMoves).find((ppm) => ppm.move_id === moveId);
  if (['DRAFT', 'SUBMITTED', 'APPROVED', 'PAYMENT_REQUESTED', 'COMPLETED'].indexOf(ppmForMove?.status) > -1) {
    return ppmForMove;
  }
  return null;
};

export const selectProGearWeightTicketAndIndexById = (state, mtoShipmentId, proGearWeightId) => {
  let proGearWeightTicket = null;
  let index = -1;
  if (proGearWeightId == null) {
    return { proGearWeightTicket, index };
  }

  const mtoShipment = selectMTOShipmentById(state, mtoShipmentId);
  const proGearWeightTickets = mtoShipment?.ppmShipment?.proGearWeightTickets;
  if (Array.isArray(proGearWeightTickets)) {
    index = proGearWeightTickets.findIndex((ele) => ele.id === proGearWeightId);
    proGearWeightTicket = proGearWeightTickets?.[index] || null;
  }
  return { proGearWeightTicket, index };
};

export const selectCurrentPPM = (state) => {
  const move = selectCurrentMove(state);
  return selectPPMForMove(state, move?.id);
};

export const selectHasCurrentPPM = (state) => {
  return !!selectCurrentPPM(state);
};

export function selectPPMEstimateRange(state) {
  return state.entities?.ppmEstimateRanges?.undefined || null;
}

export function selectPPMSitEstimate(state) {
  return state.entities?.ppmSitEstimate?.undefined?.estimate || null;
}

export function selectReimbursementById(state, reimbursementId) {
  return state.entities?.reimbursements?.[`${reimbursementId}`] || null;
}

export const selectWeightAllotmentsForLoggedInUser = createSelector(
  selectServiceMemberFromLoggedInUser,
  selectCurrentOrders,
  (serviceMember, orders) => {
    const weightAllotment = {
      pro_gear: serviceMember.weight_allotment?.pro_gear_weight,
      pro_gear_spouse: orders?.spouse_has_pro_gear ? serviceMember.weight_allotment?.pro_gear_weight_spouse : 0,
    };

    if (orders?.has_dependents) {
      weightAllotment.weight = serviceMember.weight_allotment?.total_weight_self_plus_dependents;
    } else {
      weightAllotment.weight = serviceMember.weight_allotment?.total_weight_self;
    }

    weightAllotment.sum = [weightAllotment.weight, weightAllotment.pro_gear, weightAllotment.pro_gear_spouse].reduce(
      (acc, num) => acc + num,
      0,
    );

    return weightAllotment;
  },
);

export const selectProGearEntitlements = (state) => {
  const orders = selectCurrentOrders(state);
  return orders?.entitlement || null;
};
