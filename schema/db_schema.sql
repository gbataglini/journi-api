CREATE TABLE user_details (
    id SERIAL PRIMARY KEY NOT NULL, 
    first_name VARCHAR (255) NOT NULL,  
    last_name VARCHAR (255) NOT NULL, 
    email VARCHAR (255) NOT NULL, 
    date_of_birth DATE, 
    created_at TIMESTAMP DEFAULT NOW(), 
    modified_at TIMESTAMP DEFAULT NULL
); 


CREATE TABLE destinations (
    id SERIAL PRIMARY KEY NOT NULL, 
    googlemaps_id VARCHAR (255) NOT NULL, 
    user_id INT NOT NULL REFERENCES user_details (id), 
    city VARCHAR (255) NOT NULL, 
    country VARCHAR (255) NOT NULL,
    visited BOOLEAN NOT NULL DEFAULT FALSE, 
    destination_type VARCHAR (255) NOT NULL, 
    created_at TIMESTAMP DEFAULT NOW(), 
    modified_at TIMESTAMP DEFAULT NULL
); 

CREATE TABLE journi_planner (
    id SERIAL PRIMARY KEY NOT NULL, 
    destination_id INT NOT NULL REFERENCES destinations (googlemaps_id),
    user_id INT NOT NULL REFERENCES user_details (id), 
    date_from TIMESTAMP, 
    date_to TIMESTAMP, 
    accommodation JSON, 
    trip_type VARCHAR (255), 
    created_at TIMESTAMP DEFAULT NOW(),  
    modified_at TIMESTAMP DEFAULT NULL
); 


CREATE TABLE journi_plans (
    id SERIAL PRIMARY KEY NOT NULL, 
    destination_id INT NOT NULL REFERENCES destinations (id),
    user_id INT NOT NULL REFERENCES user_details (id), 
    planner_id INT NOT NULL REFERENCES journi_planner (id), 
    location_title VARCHAR (255) NOT NULL, 
    location_address VARCHAR (255) NOT NULL, 
    location_type VARCHAR (255) NOT NULL, 
    trip_duration INT, 
    planned_visit_day INT, 
    created_at TIMESTAMP DEFAULT NOW(),  
    modified_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE journal_entries (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL REFERENCES user_details (id), 
    destination_id INT NOT NULL REFERENCES destinations (id),
    planner_id INT NOT NULL REFERENCES journi_planner (id), 
    entry_text TEXT,
    entry_location JSON, 
    media TEXT,
    created_at TIMESTAMP DEFAULT NOW(), 
    modified_at TIMESTAMP DEFAULT NULL
); 