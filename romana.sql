CREATE TABLE networks (
	id INTEGER PRIMARY KEY,   
	name VARCHAR,
	cidr VARCHAR,
	block_mask INTEGER
)

CREATE TABLE groups (
	id INTEGER PRIMARY KEY,   
	group_id INTEGER,
	routing VARCHAR,
	assignment VARCHAR,
	cidr VARCHAR,
	FOREIGN KEY (group_id) REFERENCES groups(group_id)
 )

CREATE TABLE hosts (
	ip VARCHAR,
	name VARCHAR,
	group_id INTEGER,
	FOREIGN KEY (group_id) REFERENCES groups(group_id)
)

CREATE TABLE topologies (
   network_id INTEGER,
   group_id INTEGER,
   FOREIGN KEY (network_id) REFERENCES networks(network_id),
   FOREIGN KEY (group_id) REFERENCES groups(group_id)
)

CREATE TABLE assignments (
    id INTEGER PRIMARY KEY,
    name VARCHAR,
    val VARCHAR
)

CREATE TABLE assignments_groups (
	assignment_id INTEGER,
	FOREIGN KEY (assignment_id) REFERENCES assignment(assignment_id)
	group_id INTEGER,
	FOREIGN KEY (group_id) REFERENCES groups(group_id)
)

 
CREATE TABLE blocks (
     id INTEGER PRIMARY KEY,
     group_id INTEGER,
     FOREIGN KEY (group_id) REFERENCES groups(group_id),
     owner VARCHAR,
     cidr VARCHAR,
     ring BLOB
)
      