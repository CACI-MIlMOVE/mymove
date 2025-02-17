
-- This migration allows a CAC cert to have read/write access to all orders and the prime API.
-- The Orders API and the Prime API use client certificate authentication. Only certificates
-- signed by a trusted CA (such as DISA) are allowed which includes CACs.
-- Using a person's CAC as the certificate is a convenient way to permit a
-- single trusted individual to interact with the Orders API and the Prime API. Eventually
-- this CAC certificate should be removed.
INSERT INTO public.client_certs (
    id,
    sha256_digest,
    subject,
    allow_dps_auth_api,
    allow_orders_api,
    created_at,
    updated_at,
    allow_air_force_orders_read,
    allow_air_force_orders_write,
    allow_army_orders_read,
    allow_army_orders_write,
    allow_coast_guard_orders_read,
    allow_coast_guard_orders_write,
    allow_marine_corps_orders_read,
    allow_marine_corps_orders_write,
    allow_navy_orders_read,
    allow_navy_orders_write,
    allow_prime)
VALUES (
           '9928caf4-072c-438a-8a5e-07b213bc1826',
           'c296fff33c58131dad943ff7fd519f45a9212889bf1258f3b50060eba5a74120',
           'CN=ledeep,OU=DoD+OU=PKI+OU=CONTRACTOR,O=U.S. Government,C=US',
           false,
           true,
           now(),
           now(),
           true,
           true,
           true,
           true,
           true,
           true,
           true,
           true,
           true,
           true,
           true);
