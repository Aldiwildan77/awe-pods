create type gender as enum('Man', 'Woman');

create table if not exists users (
	id uuid not null,
	username varchar(25) not null,
	email varchar(255) not null,
	password varchar(255) not null,
	name varchar(255) not null,
	age int not null,
	gender gender,
	verified bool not null default false,
	description text null,
	photo_url varchar(255) null,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint users_pkey primary key (id),
	constraint users_email_unique unique (email)
);
create index users_name_idx on users using btree (name);
create index users_age_idx on users using btree (age);

create table if not exists category (
	id uuid not null,
	category varchar(100) not null,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint category_pkey primary key (id)
);

create table if not exists podcast (
	id uuid not null,
	owner_id uuid not null,
	title varchar(255) not null,
	slug varchar(50) not null,
	description text null,
	is_public bool default false,
	rating decimal default 0,
	file_url varchar(255) null,
	thumbnail_url varchar(255) null,
  created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint podcast_pkey primary key (id),
	constraint podcast_slug_unique unique (slug),
	constraint podcast_owner_id_fkey foreign key (owner_id) references users(id) on delete restrict
);

create table if not exists users_subscription (
	id uuid not null,
	user_id uuid not null,
	subscriber_id uuid not null, 
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint users_subscription_pkey primary key (id),
	constraint users_subscription_user_id_fkey foreign key (user_id) references users(id) on delete restrict,
	constraint users_subscription_subscriber_id_fkey foreign key (subscriber_id) references users(id) on delete restrict
);

create table if not exists podcast_rating (
	id uuid not null,
	podcast_id uuid not null,
	user_id uuid not null,
	rating smallint check (rating between 1 and 5),
  created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint podcast_rating_pkey primary key (id),
	constraint podcast_rating_podcast_id_fkey foreign key (podcast_id) references podcast(id) on delete restrict,
	constraint podcast_rating_user_id_fkey foreign key (user_id) references users(id) on delete restrict
);

create table if not exists playlist (
	id uuid not null,
	title varchar(255) not null,
	slug varchar(50) not null,
	description text null,
	owner_id uuid not null,
	is_public bool default false,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint playlist_pkey primary key (id),
	constraint playlist_slug_unique unique (slug),
	constraint playlist_owner_id_fkey foreign key (owner_id) references users(id) on delete restrict
);

create table if not exists podcast_playlist (
	id uuid not null,
	playlist_id uuid not null,
	podcast_id uuid not null,
	order_number int not null, 
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint podcast_playlist_pkey primary key (id),
	constraint podcast_playlist_playlist_id_fkey foreign key (playlist_id) references playlist(id) on delete restrict,
	constraint podcast_playlist_podcast_id_fkey foreign key (podcast_id) references podcast(id) on delete restrict
);

create table if not exists podcast_users_comments (
	id uuid not null,
	comment_parent_id uuid null,
	user_id uuid not null,
	podcast_id uuid not null,
	"comments" varchar(5000) not null, 
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint podcast_users_comment_pkey primary key (id),
	constraint podcast_users_comments_comment_parent_id_fkey foreign key (comment_parent_id) references podcast_users_comments(id) on delete restrict,
	constraint podcast_users_comments_user_id_fkey foreign key (user_id) references users(id) on delete restrict,
	constraint podcast_users_comments_podcast_id_fkey foreign key (podcast_id) references podcast(id) on delete restrict
);

create table if not exists podcast_category (
	id uuid not null,
	category_id uuid not null,
	podcast_id uuid not null,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamptz NULL,
	constraint podcast_category_pkey primary key (id),
	constraint podcast_category_category_id_fkey foreign key (category_id) references category(id) on delete restrict,
	constraint podcast_category_podcast_id_fkey foreign key (podcast_id) references podcast(id) on delete restrict
);
