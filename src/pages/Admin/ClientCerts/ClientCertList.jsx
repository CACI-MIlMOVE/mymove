import React from 'react';
import { BooleanField, Datagrid, Filter, List, TextField, TextInput } from 'react-admin';
import { makeStyles } from '@material-ui/core/styles';

import AdminPagination from 'scenes/SystemAdmin/shared/AdminPagination';

const ClientCertListFilter = (props) => (
  // eslint-disable-next-line react/jsx-props-no-spreading
  <Filter {...props}>
    <TextInput labl="Search by Cert Subject" source="search" resettable alwaysOn />
  </Filter>
);

const defaultSort = { field: 'subject', order: 'ASC' };

const useStyles = makeStyles(() => ({
  tableCell: {
    maxWidth: 150,
    whiteSpace: 'normal',
    overflow: 'scroll',
    overflowWrap: 'break-word',
  },
}));

const ClientCertList = (props) => {
  const classes = useStyles();
  return (
    <List
      // eslint-disable-next-line react/jsx-props-no-spreading
      {...props}
      pagination={<AdminPagination />}
      perPage={25}
      bulkActionButtons={false}
      sort={defaultSort}
      classes={{
        dense: 'MuiDataGrid-root--densityCompact',
      }}
      dense
      filters={<ClientCertListFilter />}
    >
      <Datagrid rowClick="show">
        <TextField cellClassName={classes.tableCell} source="subject" />
        <BooleanField source="allowOrdersAPI" label="Allow Orders API" />
        <BooleanField source="allowAirForceOrdersRead" />
        <BooleanField source="allowAirForceOrdersWrite" />
        <BooleanField source="allowArmyOrdersRead" />
        <BooleanField source="allowArmyOrdersWrite" />
        <BooleanField source="allowCoastGuardOrdersRead" />
        <BooleanField source="allowCoastGuardOrdersWrite" />
        <BooleanField source="allowMarineCorpsOrdersRead" />
        <BooleanField source="allowMarineCorpsOrdersWrite" />
        <BooleanField source="allowNavyOrdersRead" />
        <BooleanField source="allowNavyOrdersWrite" />
        <BooleanField source="allowPrime" />
        <TextField source="id" />
        <TextField source="sha256Digest" />
        <TextField source="userId" label="User Id" />
      </Datagrid>
    </List>
  );
};

export default ClientCertList;
