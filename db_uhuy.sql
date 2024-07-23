-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 22 Jul 2024 pada 15.41
-- Versi server: 10.4.32-MariaDB-log
-- Versi PHP: 8.2.12
SET
  SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

START TRANSACTION;

SET
  time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;

/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_uhuy`
--
-- --------------------------------------------------------
--
-- Struktur dari tabel `tb_users`
--
CREATE TABLE
  `users` (
    `id` int (12) NOT NULL,
    `name` varchar(24) NOT NULL,
    `username` varchar(24) NOT NULL,
    `password` varchar(100) NOT NULL,
    `address` varchar(24) NOT NULL,
    `age` int (12) NOT NULL
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

--
-- Indexes for dumped tables
--
--
-- Indeks untuk tabel `tb_users`
--
ALTER TABLE `users` ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--
--
-- AUTO_INCREMENT untuk tabel `tb_users`
--
ALTER TABLE `users` MODIFY `id` int (12) NOT NULL AUTO_INCREMENT;

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;

/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- for postgresql
CREATE TABLE
  users (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY NOT NULL,
    name varchar(24) NOT NULL,
    username varchar(24) NOT NULL,
    password varchar(100) NOT NULL,
    address varchar(24) NOT NULL,
    age int NOT NULL
  );