DELETE FROM sensitive_data;
ALTER TABLE sensitive_data AUTO_INCREMENT=0;
DELETE FROM student;
ALTER TABLE student AUTO_INCREMENT=0;
DELETE FROM employee;
ALTER TABLE employee AUTO_INCREMENT=0;
DELETE FROM accession;
ALTER TABLE accession AUTO_INCREMENT=0;
DELETE FROM people;
ALTER TABLE people AUTO_INCREMENT=0;
DELETE FROM ranks;
ALTER TABLE ranks AUTO_INCREMENT=0;
DELETE FROM groups;
ALTER TABLE groups AUTO_INCREMENT=0;
DELETE FROM speciality;
ALTER TABLE speciality AUTO_INCREMENT=0;
DELETE FROM direction;
ALTER TABLE direction AUTO_INCREMENT=0;
DELETE FROM cathedra;
ALTER TABLE cathedra AUTO_INCREMENT=0;
DELETE FROM faculty;
ALTER TABLE faculty AUTO_INCREMENT=0;
DELETE FROM institute;
ALTER TABLE institute AUTO_INCREMENT=0;


INSERT INTO institute(name) VALUES ("institute1");
INSERT INTO institute(name) VALUES ("institute2");

INSERT INTO faculty(name, id_institute) VALUES ("faculty1", 1);
INSERT INTO faculty(name, id_institute) VALUES ("faculty2", 1);
INSERT INTO faculty(name, id_institute) VALUES ("faculty3", 2);
INSERT INTO faculty(name, id_institute) VALUES ("faculty4", 2);

INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra1", 1);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra2", 1);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra3", 2);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra4", 2);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra5", 3);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra6", 3);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra7", 4);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra8", 4);

INSERT INTO direction(name, id_cathedra) VALUES ("direction1", 1);
INSERT INTO direction(name, id_cathedra) VALUES ("direction2", 1);
INSERT INTO direction(name, id_cathedra) VALUES ("direction3", 2);
INSERT INTO direction(name, id_cathedra) VALUES ("direction4", 2);
INSERT INTO direction(name, id_cathedra) VALUES ("direction5", 3);
INSERT INTO direction(name, id_cathedra) VALUES ("direction6", 3);
INSERT INTO direction(name, id_cathedra) VALUES ("direction7", 4);
INSERT INTO direction(name, id_cathedra) VALUES ("direction8", 4);
INSERT INTO direction(name, id_cathedra) VALUES ("direction9", 5);
INSERT INTO direction(name, id_cathedra) VALUES ("direction10", 5);
INSERT INTO direction(name, id_cathedra) VALUES ("direction11", 6);
INSERT INTO direction(name, id_cathedra) VALUES ("direction12", 6);
INSERT INTO direction(name, id_cathedra) VALUES ("direction13", 7);
INSERT INTO direction(name, id_cathedra) VALUES ("direction14", 7);
INSERT INTO direction(name, id_cathedra) VALUES ("direction15", 8);
INSERT INTO direction(name, id_cathedra) VALUES ("direction16", 8);

INSERT INTO speciality(name, id_direction) VALUES ("speciality1", 1);
INSERT INTO speciality(name, id_direction) VALUES ("speciality2", 1);
INSERT INTO speciality(name, id_direction) VALUES ("speciality3", 2);
INSERT INTO speciality(name, id_direction) VALUES ("speciality4", 2);
INSERT INTO speciality(name, id_direction) VALUES ("speciality5", 3);
INSERT INTO speciality(name, id_direction) VALUES ("speciality6", 3);
INSERT INTO speciality(name, id_direction) VALUES ("speciality7", 4);
INSERT INTO speciality(name, id_direction) VALUES ("speciality8", 4);
INSERT INTO speciality(name, id_direction) VALUES ("speciality9", 5);
INSERT INTO speciality(name, id_direction) VALUES ("speciality10", 5);
INSERT INTO speciality(name, id_direction) VALUES ("speciality11", 6);
INSERT INTO speciality(name, id_direction) VALUES ("speciality12", 6);
INSERT INTO speciality(name, id_direction) VALUES ("speciality13", 7);
INSERT INTO speciality(name, id_direction) VALUES ("speciality14", 7);
INSERT INTO speciality(name, id_direction) VALUES ("speciality15", 8);
INSERT INTO speciality(name, id_direction) VALUES ("speciality16", 8);
INSERT INTO speciality(name, id_direction) VALUES ("speciality17", 9);
INSERT INTO speciality(name, id_direction) VALUES ("speciality18", 9);
INSERT INTO speciality(name, id_direction) VALUES ("speciality19", 10);
INSERT INTO speciality(name, id_direction) VALUES ("speciality20", 10);
INSERT INTO speciality(name, id_direction) VALUES ("speciality21", 11);
INSERT INTO speciality(name, id_direction) VALUES ("speciality22", 11);
INSERT INTO speciality(name, id_direction) VALUES ("speciality23", 12);
INSERT INTO speciality(name, id_direction) VALUES ("speciality24", 12);
INSERT INTO speciality(name, id_direction) VALUES ("speciality25", 13);
INSERT INTO speciality(name, id_direction) VALUES ("speciality26", 13);
INSERT INTO speciality(name, id_direction) VALUES ("speciality27", 14);
INSERT INTO speciality(name, id_direction) VALUES ("speciality28", 14);
INSERT INTO speciality(name, id_direction) VALUES ("speciality29", 15);
INSERT INTO speciality(name, id_direction) VALUES ("speciality30", 15);
INSERT INTO speciality(name, id_direction) VALUES ("speciality31", 16);
INSERT INTO speciality(name, id_direction) VALUES ("speciality32", 16);

INSERT INTO groups(name, id_speciality) VALUES ("group1", 1);
INSERT INTO groups(name, id_speciality) VALUES ("group2", 1);
INSERT INTO groups(name, id_speciality) VALUES ("group3", 2);
INSERT INTO groups(name, id_speciality) VALUES ("group4", 2);
INSERT INTO groups(name, id_speciality) VALUES ("group5", 3);
INSERT INTO groups(name, id_speciality) VALUES ("group6", 3);
INSERT INTO groups(name, id_speciality) VALUES ("group7", 4);
INSERT INTO groups(name, id_speciality) VALUES ("group8", 4);
INSERT INTO groups(name, id_speciality) VALUES ("group9", 5);
INSERT INTO groups(name, id_speciality) VALUES ("group10", 5);
INSERT INTO groups(name, id_speciality) VALUES ("group11", 6);
INSERT INTO groups(name, id_speciality) VALUES ("group12", 6);
INSERT INTO groups(name, id_speciality) VALUES ("group13", 7);
INSERT INTO groups(name, id_speciality) VALUES ("group14", 7);
INSERT INTO groups(name, id_speciality) VALUES ("group15", 8);
INSERT INTO groups(name, id_speciality) VALUES ("group16", 8);
INSERT INTO groups(name, id_speciality) VALUES ("group17", 9);
INSERT INTO groups(name, id_speciality) VALUES ("group18", 9);
INSERT INTO groups(name, id_speciality) VALUES ("group19", 10);
INSERT INTO groups(name, id_speciality) VALUES ("group20", 10);
INSERT INTO groups(name, id_speciality) VALUES ("group21", 11);
INSERT INTO groups(name, id_speciality) VALUES ("group22", 11);
INSERT INTO groups(name, id_speciality) VALUES ("group23", 12);
INSERT INTO groups(name, id_speciality) VALUES ("group24", 12);
INSERT INTO groups(name, id_speciality) VALUES ("group25", 13);
INSERT INTO groups(name, id_speciality) VALUES ("group26", 13);
INSERT INTO groups(name, id_speciality) VALUES ("group27", 14);
INSERT INTO groups(name, id_speciality) VALUES ("group28", 14);
INSERT INTO groups(name, id_speciality) VALUES ("group29", 15);
INSERT INTO groups(name, id_speciality) VALUES ("group30", 15);
INSERT INTO groups(name, id_speciality) VALUES ("group31", 16);
INSERT INTO groups(name, id_speciality) VALUES ("group32", 16);
INSERT INTO groups(name, id_speciality) VALUES ("group33", 17);
INSERT INTO groups(name, id_speciality) VALUES ("group34", 17);
INSERT INTO groups(name, id_speciality) VALUES ("group35", 18);
INSERT INTO groups(name, id_speciality) VALUES ("group36", 18);
INSERT INTO groups(name, id_speciality) VALUES ("group37", 19);
INSERT INTO groups(name, id_speciality) VALUES ("group38", 19);
INSERT INTO groups(name, id_speciality) VALUES ("group39", 20);
INSERT INTO groups(name, id_speciality) VALUES ("group40", 20);
INSERT INTO groups(name, id_speciality) VALUES ("group41", 21);
INSERT INTO groups(name, id_speciality) VALUES ("group42", 21);
INSERT INTO groups(name, id_speciality) VALUES ("group43", 22);
INSERT INTO groups(name, id_speciality) VALUES ("group44", 22);
INSERT INTO groups(name, id_speciality) VALUES ("group45", 23);
INSERT INTO groups(name, id_speciality) VALUES ("group46", 23);
INSERT INTO groups(name, id_speciality) VALUES ("group47", 24);
INSERT INTO groups(name, id_speciality) VALUES ("group48", 24);
INSERT INTO groups(name, id_speciality) VALUES ("group49", 25);
INSERT INTO groups(name, id_speciality) VALUES ("group50", 25);
INSERT INTO groups(name, id_speciality) VALUES ("group51", 26);
INSERT INTO groups(name, id_speciality) VALUES ("group52", 26);
INSERT INTO groups(name, id_speciality) VALUES ("group53", 27);
INSERT INTO groups(name, id_speciality) VALUES ("group54", 27);
INSERT INTO groups(name, id_speciality) VALUES ("group55", 28);
INSERT INTO groups(name, id_speciality) VALUES ("group56", 28);
INSERT INTO groups(name, id_speciality) VALUES ("group57", 29);
INSERT INTO groups(name, id_speciality) VALUES ("group58", 29);
INSERT INTO groups(name, id_speciality) VALUES ("group59", 30);
INSERT INTO groups(name, id_speciality) VALUES ("group60", 30);
INSERT INTO groups(name, id_speciality) VALUES ("group61", 31);
INSERT INTO groups(name, id_speciality) VALUES ("group62", 31);
INSERT INTO groups(name, id_speciality) VALUES ("group63", 32);
INSERT INTO groups(name, id_speciality) VALUES ("group64", 32);

INSERT INTO ranks(name) VALUES ("rank1");
INSERT INTO ranks(name) VALUES ("rank2");
INSERT INTO ranks(name) VALUES ("rank3");
INSERT INTO ranks(name) VALUES ("rank4");
INSERT INTO ranks(name) VALUES ("rank5");
INSERT INTO ranks(name) VALUES ("rank6");

INSERT INTO people(fio, birthday, gender, comment, password, phone_number, email, status, have_access)
	VALUES ("Степул Артем Мартиросовчи", "1997-06-12", 1, "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "artem.stepul@gmail.com", "учится", 1);
INSERT INTO people(fio, birthday, gender, comment, password, phone_number, email, status, have_access)
	VALUES ("Луценко Артем Геннадиевич", "1997-03-16", 1, "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "artem.lucenko@gmail.com", "учится", 1);
INSERT INTO people(fio, birthday, gender, comment, password, phone_number, email, status, have_access)
	VALUES ("Иванов Иван Иванович", "1979-09-17", 1, "", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "ivanov@gmail.com", "работает", 1);

INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
	VALUES (1, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
	VALUES (2, "2017-09-01", 1, 1, 4, 1);

INSERT INTO employee(id_people, date_invite, id_group, id_cathedra, id_rank)
	VALUES (3, "2007-07-22", NULL, 1, 1);
	
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
	VALUES (1, 0, 0, 0, 0, 0, 0, 0, 0, 0);	
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
	VALUES (2, 0, 0, 0, 0, 0, 0, 0, 0, 0);	
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
	VALUES (3, 1, 1, 1, 1, 1, 1, 1, 1, 1);

INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
	VALUES (1, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
	VALUES (2, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
	VALUES (3, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");