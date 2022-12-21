USE tokoku;

SHOW tables;

DROP TABLE item;
DROP TABLE transaksi;
DROP TABLE barang;
DROP TABLE konsumen;
DROP TABLE pegawai;

CREATE TABLE pegawai (
	id_pegawai int AUTO_INCREMENT NOT NULL,
	nama varchar(50) NOT NULL,
	`password` varchar(50) NOT NULL,
	PRIMARY KEY (id_pegawai)
);

CREATE TABLE konsumen (
	hp_konsumen varchar(13) NOT NULL,
	nama varchar(50) NOT NULL,
	id_pegawai int NOT NULL,
	CONSTRAINT fk_id_pegawai FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai),
	PRIMARY KEY (hp_konsumen)
);

CREATE TABLE barang (
	id_barang int AUTO_INCREMENT NOT NULL,
	nama_barang varchar(50) NOT NULL,
	deskripsi varchar(50) NOT NULL,
	stok int NOT NULL,
	id_pegawai int NOT NULL,
	FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai),
	PRIMARY KEY (id_barang)
);

CREATE TABLE transaksi (
	no_nota int AUTO_INCREMENT NOT NULL,
	id_pegawai int NOT NULL,
	hp_konsumen varchar(13) NOT NULL,
	tanggal_cetak timestamp DEFAULT now(),
	FOREIGN KEY (id_pegawai) REFERENCES pegawai(id_pegawai),
	FOREIGN KEY (hp_konsumen) REFERENCES konsumen(hp_konsumen),
	PRIMARY KEY (no_nota)
);

CREATE TABLE item (
	no_nota int NOT NULL,
	id_barang int NOT NULL,
	kuantitas int NOT NULL,
	FOREIGN KEY (no_nota) REFERENCES transaksi(no_nota),
	FOREIGN KEY (id_barang) REFERENCES barang(id_barang),
	PRIMARY KEY (no_nota, id_barang)
);

INSERT INTO pegawai VALUES (NULL,'admin','admin');
INSERT INTO pegawai VALUES (NULL,'andi','andi123');

INSERT INTO konsumen VALUES ('08123456789','bejo',1);
INSERT INTO konsumen VALUES ('08111','budi',2);

INSERT INTO barang VALUES (NULL,'indomie','mie goreng', 48, 1);
INSERT INTO barang VALUES (NULL,'le minerale','air minum', 20, 1);

INSERT INTO transaksi VALUES (NULL, 1, '08123456789', now());

INSERT INTO item VALUES (1, 1, 5);

SELECT * FROM pegawai p ;
SELECT * FROM konsumen k ;
SELECT * FROM barang b ;
SELECT * FROM transaksi t ;
SELECT * FROM item i ;