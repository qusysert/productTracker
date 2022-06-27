package repository

import (
	"database/sql"
)

func Migrate(conn *sql.DB) error {
	list := []string{
		`create table if not exists seller
(
    id int not null
        constraint seller_pk
            primary key
)`,
		`create table if not exists product
(
    id        serial
        constraint product_pk
            primary key,
    seller_id integer not null
        constraint product_seller_id_fk
            references seller
            on delete cascade,
    offer_id  integer not null,
    name      text    not null,
    price     integer not null,
    quantity  integer not null,
    available boolean not null
)`,

		`create index if not exists product_name_index
    on product (name)`,

		`create index if not exists product_offer_id_index
    on product (offer_id)`,

		`create index if not exists product_seller_id_index
    on product (seller_id)`,
	}

	for _, q := range list {
		_, err := conn.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil
}
