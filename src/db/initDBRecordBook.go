package db

const createTableInstitute = `CREATE TABLE IF NOT EXISTS institute(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableFaculty = `CREATE TABLE IF NOT EXISTS faculty(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL,
										id_institute INT NOT NULL,
										FOREIGN KEY	(id_institute) REFERENCES institute(id) ON DELETE RESTRICT
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableCathedra = `CREATE TABLE IF NOT EXISTS cathedra(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL,
										id_faculty INT NOT NULL,
										FOREIGN KEY	(id_faculty) REFERENCES faculty(id) ON DELETE RESTRICT
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableDirection = `CREATE TABLE IF NOT EXISTS direction(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL, 
										id_cathedra INT NOT NULL,
										FOREIGN KEY	(id_cathedra) REFERENCES cathedra(id) ON DELETE RESTRICT
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableSpeciality = `CREATE TABLE IF NOT EXISTS speciality(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL, 
										id_direction INT NOT NULL,
										FOREIGN KEY	(id_direction) REFERENCES direction(id) ON DELETE RESTRICT
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableGroups = `CREATE TABLE IF NOT EXISTS groups(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL,
										id_speciality INT,
										FOREIGN KEY (id_speciality) REFERENCES speciality(id) ON DELETE RESTRICT ON UPDATE CASCADE
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableRanks = `CREATE TABLE IF NOT EXISTS ranks(
										id INT PRIMARY KEY AUTO_INCREMENT,
										name VARCHAR(64) NOT NULL
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTablePeople = `CREATE TABLE IF NOT EXISTS people(
										id INT PRIMARY KEY AUTO_INCREMENT,
										fio VARCHAR(256) NOT NULL,
										birthday DATE DEFAULT NULL,
										sex TINYINT(1) NOT NULL,
										comment TEXT,
										password VARCHAR(256) NOT NULL,
										phone_number VARCHAR(64),
										email VARCHAR(64) NOT NULL UNIQUE,
										status VARCHAR(256) NOT NULL,
										have_access TINYINT(1) NOT NULL DEFAULT 1
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableStudent = `CREATE TABLE IF NOT EXISTS student(
										id_people INT PRIMARY KEY REFERENCES people(id),
										date_admission DATE NOT NULL,
										is_full_time TINYINT(1) NOT NULL,
										is_cut TINYINT(1) NOT NULL,
										semester INT NOT NULL DEFAULT 1,
										id_group INT,
										FOREIGN KEY (id_group) REFERENCES groups(id) ON DELETE RESTRICT ON UPDATE CASCADE
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableEmployee = `CREATE TABLE IF NOT EXISTS employee(
										id_people INT PRIMARY KEY REFERENCES people(id),
										date_invite DATE NOT NULL,
										id_group INT DEFAULT NULL,
										id_cathedra INT NOT NULL,
										id_rank INT NOT NULL,
										FOREIGN KEY (id_rank) REFERENCES ranks(id) ON DELETE RESTRICT ON UPDATE CASCADE,
										FOREIGN KEY (id_group) REFERENCES groups(id) ON DELETE RESTRICT ON UPDATE CASCADE,
										FOREIGN KEY (id_cathedra) REFERENCES cathedra(id) ON DELETE RESTRICT ON UPDATE CASCADE
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableAccession = `CREATE TABLE IF NOT EXISTS accession(
										id_people INT PRIMARY KEY REFERENCES people(id),
										edit_access TINYINT(1) NOT NULL DEFAULT 0,
										set_absence TINYINT(1) NOT NULL DEFAULT 0,
										get_absence TINYINT(1) NOT NULL DEFAULT 0,
										set_mark TINYINT(1) NOT NULL DEFAULT 0,
										set_event TINYINT(1) NOT NULL DEFAULT 0,
										get_sensitive TINYINT(1) NOT NULL DEFAULT 0,
										set_sensitive TINYINT(1) NOT NULL DEFAULT 0,
										get_ylist TINYINT(1) NOT NULL DEFAULT 0,
										manage_academ TINYINT(1) NOT NULL DEFAULT 0
								) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

const createTableSensitiveData = `CREATE TABLE IF NOT EXISTS sensitive_data(
										id_people INT PRIMARY KEY REFERENCES people(id) ON DELETE CASCADE,
										passport_code VARCHAR(512) DEFAULT NULL,
										rntrs VARCHAR(512) DEFAULT NULL,
										reg_address VARCHAR(512) DEFAULT NULL,
										military_id VARCHAR(512) DEFAULT NULL
									) ENGINE=InnoDB DEFAULT CHARSET="utf8";`

var queriesForInitDb = [...]string{
	createTableInstitute,
	createTableFaculty,
	createTableCathedra,
	createTableDirection,
	createTableSpeciality,
	createTableRanks,
	createTableGroups,
	createTablePeople,
	createTableStudent,
	createTableEmployee,
	createTableAccession,
	createTableSensitiveData,
}
