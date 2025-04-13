drop table if exists authorized_person cascade;
create table authorized_person
(
    birth_date   bigint       not null,
    id           bigint generated by default as identity
        primary key,
    first_name   varchar(255) not null,
    last_name    varchar(255) not null,
    phone_number varchar(255) not null
);


drop table if exists company cascade;
create table company
(
    id             bigint generated by default as identity
        primary key,
    address        varchar(255),
    company_number varchar(255),
    name           varchar(255),
    vat_number     varchar(255),
    bas            varchar(255),
    ownerid        bigint
);


drop table if exists account cascade;
create table account
(
    balance                 double precision not null,
    daily_limit             double precision not null,
    daily_spent             double precision not null,
    monthly_limit           double precision not null,
    monthly_maintenance_fee double precision not null,
    monthly_spent           double precision not null,
    reserved_balance        double precision not null,
    company_id              bigint
        constraint fkyx9qhair86pc6tnbh322xjg4
            references company,
    created_date            bigint           not null,
    employeeid              bigint           not null,
    expiration_date         bigint           not null,
    id                      bigint generated by default as identity
        primary key,
    ownerid                 bigint           not null,
    account_number          varchar(255)     not null
        unique,
    currency_type           varchar(255)     not null
        constraint account_currency_type_check
            check ((currency_type)::text = ANY
        ((ARRAY ['RSD':: character varying, 'EUR':: character varying, 'USD':: character varying, 'CHF':: character varying, 'GBP':: character varying, 'JPY':: character varying, 'CAD':: character varying, 'AUD':: character varying])::text[])
) ,
    status                  varchar(255)     not null
        constraint account_status_check
            check ((status)::text = ANY
                   ((ARRAY ['ACTIVE'::character varying, 'BLOCKED'::character varying, 'CLOSED'::character varying, 'FROZEN'::character varying])::text[])),
    subtype                 varchar(255)     not null
        constraint account_subtype_check
            check ((subtype)::text = ANY
                   ((ARRAY ['STANDARD'::character varying, 'SAVINGS'::character varying, 'PENSION'::character varying, 'STUDENT'::character varying, 'YOUTH'::character varying, 'PERSONAL'::character varying, 'BUSINESS'::character varying])::text[])),
    type                    varchar(255)     not null
        constraint account_type_check
            check ((type)::text = ANY
                   ((ARRAY ['CURRENT'::character varying, 'FOREIGN_CURRENCY'::character varying, 'BANK'::character varying, 'COUNTRY'::character varying])::text[]))
);


drop table if exists card cascade;
create table card
(
    active               boolean          not null,
    blocked              boolean          not null,
    card_limit           double precision not null,
    account_id           bigint           not null
        constraint fk8v67eys6tqflsm6hrdgru2phu
            references account,
    authorized_person_id bigint
        constraint fkxq19frv57038g5v8jed9nrym
            references authorized_person,
    created_at           bigint           not null,
    expiration_date      bigint           not null,
    id                   bigint generated by default as identity
        primary key,
    card_brand           varchar(255)     not null
        constraint card_card_brand_check
            check ((card_brand)::text = ANY
        ((ARRAY ['VISA':: character varying, 'MASTERCARD':: character varying, 'DINA_CARD':: character varying, 'AMERICAN_EXPRESS':: character varying])::text[])
) ,
    card_cvv             varchar(255)     not null,
    card_name            varchar(255)     not null,
    card_number          varchar(255)     not null,
    card_type            varchar(255)     not null
        constraint card_card_type_check
            check ((card_type)::text = ANY ((ARRAY ['DEBIT'::character varying, 'CREDIT'::character varying])::text[]))
);


drop table if exists currency cascade;
create table currency
(
    id   bigint generated by default as identity
        primary key,
    code varchar(255) not null
        unique
        constraint currency_code_check
            check ((code)::text = ANY
        ((ARRAY ['RSD':: character varying, 'EUR':: character varying, 'USD':: character varying, 'CHF':: character varying, 'GBP':: character varying, 'JPY':: character varying, 'CAD':: character varying, 'AUD':: character varying])::text[])
) ,
    country varchar(255) not null,
    name    varchar(255) not null,
    symbol  varchar(255) not null
);


drop table if exists exchange_pair cascade;
create table exchange_pair
(
    date               date             not null,
    exchange_rate      double precision not null,
    base_currency_id   bigint           not null
        constraint fkkg59kymnfgyny7rsqufplwun4
            references currency,
    id                 bigint generated by default as identity
        primary key,
    target_currency_id bigint           not null
        constraint fkef7bcvui45w4576jmvspq8u9h
            references currency
);


drop table if exists loan cascade;
create table loan
(
    effective_rate              double precision not null,
    loan_amount                 double precision not null,
    monthly_payment             double precision not null,
    nominal_rate                double precision not null,
    penalty                     double precision not null DEFAULT 0,
    number_of_installments      integer          not null,
    number_of_paid_installments integer          not null DEFAULT 0,
    remaining_amount            double precision not null,
    account_id                  bigint           not null
        constraint fkgv9cgsh4k76wmaf83ktoekpub
            references account,
    allowed_date                bigint,
    created_date                bigint           not null,
    id                          bigint generated by default as identity
        primary key,
    next_payment_date           varchar(255)     not null,
    phone_number                varchar(255)     not null,
    currency_type               varchar(255)     not null
        constraint loan_currency_type_check
            check ((currency_type)::text = ANY
        ((ARRAY ['RSD':: character varying, 'EUR':: character varying, 'USD':: character varying, 'CHF':: character varying, 'GBP':: character varying, 'JPY':: character varying, 'CAD':: character varying, 'AUD':: character varying])::text[])
) ,
    interest_type          varchar(255)     not null
        constraint loan_interest_type_check
            check ((interest_type)::text = ANY
                   ((ARRAY ['FIXED'::character varying, 'VARIABLE'::character varying])::text[])),
    loan_reason            varchar(255),
    loan_type              varchar(255)     not null
        constraint loan_loan_type_check
            check ((loan_type)::text = ANY
                   ((ARRAY ['CASH'::character varying, 'MORTGAGE'::character varying, 'AUTO'::character varying, 'REFINANCING'::character varying, 'STUDENT'::character varying])::text[])),
    payment_status         varchar(255)     not null
        constraint loan_payment_status_check
            check ((payment_status)::text = ANY
                   ((ARRAY ['PENDING'::character varying, 'APPROVED'::character varying, 'DENIED'::character varying, 'PAID_OFF'::character varying, 'LATE'::character varying])::text[]))
);


drop table if exists otp_token cascade;
create table otp_token
(
    used            boolean      not null,
    expiration_time bigint       not null,
    id              bigint generated by default as identity
        primary key,
    transfer_id     bigint       not null,
    otp_code        varchar(255) not null
);


drop table if exists receiver cascade;
create table receiver
(
    id             bigint generated by default as identity
        primary key,
    customer_id    bigint       not null, -- povezivanje sa korisnikom
    account_number varchar(255) not null,
    address        varchar(255),
    first_name     varchar(255),
    last_name      varchar(255),
    usage_count     integer default 0 -- za praćenje koliko puta je korišćen
);


drop table if exists transfer cascade;
create table transfer
(
    amount              double precision not null,
    completed_at        bigint,
    created_at          bigint,
    from_account_id     bigint           not null
        constraint fkfbrtpigqywdcc6e64ichki73j
            references account,
    from_currency_id    bigint           not null
        constraint fkc6agh3nies7d6f97aapxdhg06
            references currency,
    id                  bigint generated by default as identity
        primary key,
    to_account_id       bigint           not null
        constraint fkrppn0djf9uip7gj308swxq54j
            references account,
    to_currency_id      bigint           not null
        constraint fkbfidyemr02sgnep1ubdyj3dda
            references currency,
    adress              varchar(255),
    note                varchar(255),
    otp                 varchar(255),
    payment_code        varchar(255),
    payment_description varchar(255),
    payment_reference   varchar(255),
    receiver            varchar(255),
    status              varchar(255),
    saved_receiver_id   bigint
        constraint transfer_status_check
            check ((status)::text = ANY
        ((ARRAY ['PENDING':: character varying, 'COMPLETED':: character varying, 'CANCELLED':: character varying, 'FAILED':: character varying])::text[])
) ,
    type                varchar(255)
        constraint transfer_type_check
            check ((type)::text = ANY
                   ((ARRAY ['INTERNAL'::character varying, 'EXTERNAL'::character varying, 'EXCHANGE'::character varying, 'FOREIGN'::character varying])::text[]))
);


drop table if exists transaction cascade;
create table transaction
(
    amount          double precision not null,
    final_amount    double precision not null,
    fee             double precision not null,
    bank_only       boolean          not null,
    currency_id     bigint           not null
        constraint fklcx7g8g7x4fyns9k6vesu3n9n
            references currency,
    from_account_id bigint           not null
        constraint fkrff4jlxetafju1e5cks5mfcnk
            references account,
    id              bigint generated by default as identity
        primary key,
    timestamp       bigint           not null,
    to_account_id   bigint           not null
        constraint fkluqt8k2pa8d4gmggx4rhl5vgv
            references account,
    transfer_id     bigint           not null
        constraint fk219909wacaej3s32ttmue0irq
            references transfer,
    description     varchar(255)
);


drop table if exists installment cascade;
create table installment
(
    amount             double precision not null,
    attempt_count      integer          not null,
    installment_number integer          not null,
    currency_type      smallint         not null
        constraint installment_currency_type_check
            check ((currency_type >= 0) AND (currency_type <= 7)),
    interest_rate      double precision not null,
    is_paid            boolean          not null,
    lawsuit            boolean          not null DEFAULT FALSE,
    actual_due_date    bigint,
    expected_due_date  varchar(255)     not null,
    id                 bigint generated by default as identity
        primary key,
    loan_id            bigint           not null
        constraint fkddvr1rongdlfl3pmj87eg48cy
            references loan,
    retry_date         varchar(255),
    transaction_id     bigint
        constraint fkq04sunhaybag8qcjer2nmbhlh
            references transaction,
    payment_status     varchar(255)     not null
        constraint installment_payment_status_check
            check ((payment_status)::text = ANY
        ((ARRAY ['PENDING':: character varying, 'APPROVED':: character varying, 'DENIED':: character varying, 'PAID_OFF':: character varying, 'LATE':: character varying])::text[])
)
    );


drop table if exists rate_change cascade;
create table rate_change
(
    id     bigint generated by default as identity
        primary key,
    year   integer          not null DEFAULT 0,
    month  integer          not null DEFAULT 0,
    change double precision not null
);
