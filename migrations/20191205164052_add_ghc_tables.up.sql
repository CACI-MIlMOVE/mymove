CREATE TABLE entitlements
(
	id uuid PRIMARY KEY NOT NULL,
	dependents_authorized bool,
	total_dependents integer,
	non_temporary_storage bool,
	privately_owned_vehicle bool,
	pro_gear_weight integer,
	pro_gear_weight_spouse integer,
	storage_in_transit integer,
	created_at timestamp WITH TIME ZONE NOT NULL,
	updated_at timestamp WITH TIME ZONE NOT NULL
);

CREATE TABLE move_orders
(
	id uuid PRIMARY KEY NOT NULL,
	customer_id uuid REFERENCES customers,
	origin_duty_station_id uuid REFERENCES duty_stations,
	destination_duty_station_id uuid REFERENCES duty_stations,
	entitlement_id uuid REFERENCES entitlements,
	created_at timestamp WITH TIME ZONE NOT NULL,
	updated_at timestamp WITH TIME ZONE NOT NULL
);

CREATE TABLE move_task_orders
(
	id uuid PRIMARY KEY NOT NULL,
	move_order_id uuid REFERENCES move_orders,
	reference_id varchar,
	status ghc_approval_status NOT NULL,
	is_available_to_prime bool NOT NULL,
	is_cancelled bool NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL,
	updated_at timestamp WITH TIME ZONE NOT NULL
);

CREATE TABLE mto_shipments
(
	id uuid PRIMARY KEY NOT NULL,
	move_task_order_id uuid REFERENCES move_task_orders,
	scheduled_pickup_date date,
	requested_pickup_date date,
	customer_remarks text,
	pickup_address_id uuid REFERENCES addresses,
	destination_address_id uuid REFERENCES addresses,
	secondary_pickup_address_id uuid REFERENCES addresses,
	secondary_delivery_address_id uuid REFERENCES addresses,
	prime_estimated_weight integer,
	prime_estimated_weight_recorded_date date,
	prime_actual_weight integer,
	created_at timestamp WITH TIME ZONE NOT NULL,
	updated_at timestamp WITH TIME ZONE NOT NULL
);

CREATE TABLE mto_service_items
(
	id uuid PRIMARY KEY NOT NULL,
	move_task_order_id uuid REFERENCES move_task_orders,
	mto_shipment_id uuid REFERENCES mto_shipments,
	re_service_id uuid REFERENCES re_services,
	meta_id uuid NOT NULL,
	meta_type text NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL,
	updated_at timestamp WITH TIME ZONE NOT NULL
);

ALTER TABLE payment_requests
	ADD COLUMN move_task_order_id uuid REFERENCES move_task_orders;

ALTER TABLE customers
	ADD COLUMN first_name text NOT NULL,
	ADD COLUMN last_name text NOT NULL,
	ADD COLUMN email text NOT NULL,
	ADD COLUMN phone text NOT NULL,
	ADD COLUMN dod_id text;

CREATE INDEX ON mto_service_items (move_task_order_id);
CREATE INDEX ON mto_service_items (mto_shipment_id);
CREATE INDEX ON mto_service_items (re_service_id);

CREATE INDEX ON mto_shipments (move_task_order_id);
CREATE INDEX ON mto_shipments (pickup_address_id);
CREATE INDEX ON mto_shipments (destination_address_id);
CREATE INDEX ON mto_shipments (secondary_pickup_address_id);
CREATE INDEX ON mto_shipments (secondary_delivery_address_id);

CREATE INDEX ON move_task_orders (move_order_id);

CREATE INDEX ON move_orders (customer_id);
CREATE INDEX ON move_orders (origin_duty_station_id);
CREATE INDEX ON move_orders (destination_duty_station_id);
CREATE INDEX ON move_orders (entitlement_id);