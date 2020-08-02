package storage

const schema = `
CREATE TABLE IF NOT EXISTS users(
    id BIGINT AUTO_INCREMENT,
    name VARCHAR (100),
    country_code TINYINT (6),
    phone BIGINT NOT NULL,
    hash_pwd TEXT,
    PRIMARY KEY (id),
    UNIQUE (country_code, phone)
);

CREATE TABLE IF NOT EXISTS cabs(
    id BIGINT AUTO_INCREMENT,
    veh_no VARCHAR (20) NOT NULL,
    lat DOUBLE (10, 7),
    lon DOUBLE (10, 7),
    loc_timestamp INT,  
    driver_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (veh_no),
    FOREIGN KEY (driver_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS bookings(
    id BIGINT AUTO_INCREMENT,
    from_lat DOUBLE (10,7) NOT NULL,
    from_lon DOUBLE (10,7) NOT NULL,
    to_lat DOUBLE (10,7) NOT NULL,
    to_lon DOUBLE (10,7) NOT NULL,
    pickup_time INT NOT NULL,
    is_confirmed TINYINT (1) DEFAULT 0, 
    user_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);`
