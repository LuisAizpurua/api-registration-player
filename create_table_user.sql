create table public.users (
  id uuid not null default gen_random_uuid (),
  created_at timestamp with time zone not null default now(),
  name text not null,
  phone text not null,
  email text not null,
  session_id text null,
  constraint users_pkey primary key (id, created_at, name, phone, email)
) TABLESPACE pg_default;
