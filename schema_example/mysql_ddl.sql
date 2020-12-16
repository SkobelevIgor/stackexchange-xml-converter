-- Example of schema 
-- The schema does not include any indexes ( except PK's )

-- Users
CREATE TABLE users (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	account_id INTEGER,
	reputation INTEGER NOT NULL,
	views INTEGER DEFAULT 0,
	down_votes INTEGER DEFAULT 0,
	up_votes INTEGER DEFAULT 0,
	display_name VARCHAR(255) NOT NULL,
	location VARCHAR(512),
	profile_image_url VARCHAR(255),
	website_url VARCHAR(255),
	about_me MEDIUMTEXT,
	creation_date TIMESTAMP(3) NOT NULL,
	last_access_date TIMESTAMP(3) NOT NULL
);

-- Posts
CREATE TABLE posts (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	owner_user_id INTEGER,
	last_editor_user_id INTEGER,
	post_type_id SMALLINT NOT NULL,
	accepted_answer_id INTEGER,
	score INTEGER NOT NULL,
	parent_id INTEGER,
	view_count INTEGER,
	answer_count INTEGER DEFAULT 0,
	comment_count INTEGER DEFAULT 0,
	owner_display_name VARCHAR(64),
	last_editor_display_name VARCHAR(64),
	title VARCHAR(512),
	tags VARCHAR(512),
	content_license VARCHAR(64) NOT NULL,
	body MEDIUMTEXT,
	favorite_count INTEGER,
	creation_date TIMESTAMP(3) NOT NULL,
	community_owned_date TIMESTAMP(3),
	closed_date TIMESTAMP(3),
	last_edit_date TIMESTAMP(3),
	last_activity_date TIMESTAMP(3),
    
);

-- PostLinks
CREATE TABLE post_links (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	related_post_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	link_type_id TINYINT NOT NULL,
	creation_date TIMESTAMP(3) NOT NULL
);

-- PostHistory
CREATE TABLE post_history (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER,
	post_history_type_id TINYINT NOT NULL,
	user_display_name VARCHAR(64),
	content_license VARCHAR(64),
	revision_guid uuid,
	text MEDIUMTEXT,
	comment MEDIUMTEXT,
	creation_date TIMESTAMP(3) NOT NULL
);

-- Comments
CREATE TABLE comments (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER,
	score TINYINT NOT NULL,
	content_license VARCHAR(64) NOT NULL,
	user_display_name VARCHAR(64),
	text MEDIUMTEXT,
	creation_date TIMESTAMP(3) NOT NULL
);

-- Votes
CREATE TABLE votes (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id INTEGER,
	post_id INTEGER NOT NULL,
	vote_type_id TINYINT NOT NULL,
	bounty_amount TINYINT,
	creation_date TIMESTAMP(3) NOT NULL
);

-- Badges
CREATE TABLE badges (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id INTEGER NOT NULL,
	class TINYINT NOT NULL,
	name VARCHAR(64) NOT NULL,
	tag_based TINYINT(1) NOT NULL,
	date TIMESTAMP(3) NOT NULL
);

-- Tags
CREATE TABLE tags (
	id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
	excerpt_post_id INTEGER,
	wiki_post_id INTEGER,
	tag_name VARCHAR(255) NOT NULL,
	count INTEGER DEFAULT 0
);