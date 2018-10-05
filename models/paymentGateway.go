/**
  payment_gateway              //PG사 정보 관리 테이블
      pg_id                       // unique, not auto increse
      pg_description              //
      pg_kind                     // 향후 사용 할 수 있다. ex) 1: credit card. 2: mobile ....
      created_at                  // .defaultTo(knex.fn.now());
      closed_at                   //
*/
/*
exports.up = function(knex, Promise) {
    return Promise.all([
        knex.schema.createTable('payment_gateway', function(table) {
            table
                .integer('pg_id')
                .unsigned()
                .primary();
            table.string('pg_description').notNullable();
            table.integer('pg_kind').unsigned();
            table.timestamp('created_at').defaultTo(knex.fn.now());
            table.timestamp('closed_at');
        })
    ]);
};
*/
package models
