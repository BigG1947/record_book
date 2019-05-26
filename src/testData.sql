DELETE
from marks;
ALTER TABLE marks
    AUTO_INCREMENT = 0;
DELETE
FROM loads;
ALTER TABLE loads
    AUTO_INCREMENT = 0;
DELETE
FROM discipline;
ALTER TABLE discipline
    AUTO_INCREMENT = 0;
# DELETE FROM sensitive_data;
# ALTER TABLE sensitive_data AUTO_INCREMENT=0;
DELETE
FROM student;
ALTER TABLE student
    AUTO_INCREMENT = 0;
DELETE
FROM employee;
ALTER TABLE employee
    AUTO_INCREMENT = 0;
DELETE
FROM accession;
ALTER TABLE accession
    AUTO_INCREMENT = 0;
DELETE
FROM people;
ALTER TABLE people
    AUTO_INCREMENT = 0;
DELETE
FROM speciality;
ALTER TABLE speciality
    AUTO_INCREMENT = 0;
DELETE
FROM direction;
ALTER TABLE direction
    AUTO_INCREMENT = 0;
DELETE
FROM cathedra;
ALTER TABLE cathedra
    AUTO_INCREMENT = 0;
DELETE
FROM faculty;
ALTER TABLE faculty
    AUTO_INCREMENT = 0;
DELETE
FROM institute;
ALTER TABLE institute
    AUTO_INCREMENT = 0;
-- DELETE FROM ranks;
-- ALTER TABLE ranks AUTO_INCREMENT=0;
DELETE
FROM groups;
ALTER TABLE groups
    AUTO_INCREMENT = 0;
DELETE
FROM loads_semester;
ALTER TABLE loads_semester
    AUTO_INCREMENT = 0;
-- DELETE FROM status;
-- ALTER TABLE status AUTO_INCREMENT=0;


INSERT INTO institute(name)
VALUES ("Навчально-науковий технологічний інститут харчової промисловості ім. М.В. Ломоносова");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут комп`ютерних систем і технологій \"Індустрія 4.0\" ім. П.М. Платонова");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут прикладної економіки та менеджменту ім. Г.Е. Вейнштейна");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут холоду, кріотехнологій та екоенергетики ім. Мартиновського В.С.");


INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології зерна і зернового бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет інноваційних технологій харчування і ресторанно-готельного бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології та товарознавства харчових продуктів і продовольчого бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології вина та туристичного бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет комп`ютерних систем та автоматизації", 2);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет економіки, бізнесу і контролю", 3);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет менеджменту, маркетингу і логістики", 3);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет низькотемпературної техніки та інженерної механіки", 4);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет нафти, газу та екології", 4);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет комп`ютерної інженерії, програмування та кіберзахисту", 2);


INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології переробки зерна", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології зберігання зерна", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технологічного обладнання зернових виробництв", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології комбікормів і біопалива", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра управління бізнесом", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології хліба, кондитерських, макаронних виробів і харчоконцентратів", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології ресторанного і оздоровчого харчування", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра готельно-ресторанного бізнесу", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра інженерної графіки та технічного дизайну", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра безпеки життєдіяльності", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра іноземних мов", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології м´яса, риби і морепродуктів", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології молочних, олійно-жирових продуктів і косметики", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра харчової хімії та експертизи", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра економічної теорії та фінансово-економічної безпеки", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра товарознавства та митної справи", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології вина та енології", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра біохімії, мікробіології та фізіології харчування", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра біоінженерії і води", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра туристичного бізнесу та рекреації", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра автоматизації технологічних процесів і робототехнічних систем", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра електромеханіки та мехатроніки", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра вищої та прикладної математики", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра обліку і аудиту", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра економіки промисловості", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра українознавства і лінгводидактики", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра менеджменту і логістики", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра маркетингу, підприємництва і торгівлі", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра соціології, філософії і права", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра компресорів та пневмоагрегатів", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра кріогенної техніки", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра холодильних установок і кондиціювання повітря", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра процесів, обладнання та енергетичного менеджменту", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра фізики i матеріалознавства", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра термодинаміки та відновлювальної енергетики", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра теплоенергетики і трубопровідного транспорту енергоносіїв", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра екології та природоохоронні технології", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра теплофізики та прикладної екології", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра комп`ютерної інженерії", 10);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра інформаційних технологій та кібербезпеки", 10);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра фізичної культури та спорту", 10);

INSERT INTO direction(name, id_cathedra)
VALUES ("Комп''ютерні науки", 40);
INSERT INTO direction(name, id_cathedra)
VALUES ("Комп''ютерна інжеренерія", 40);

# INSERT INTO speciality(name, id_direction) VALUES ("speciality1", 1);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality2", 1);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality3", 2);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality4", 2);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality5", 3);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality6", 3);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality7", 4);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality8", 4);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality9", 5);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality10", 5);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality11", 6);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality12", 6);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality13", 7);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality14", 7);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality15", 8);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality16", 8);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality17", 9);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality18", 9);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality19", 10);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality20", 10);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality21", 11);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality22", 11);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality23", 12);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality24", 12);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality25", 13);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality26", 13);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality27", 14);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality28", 14);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality29", 15);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality30", 15);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality31", 16);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality32", 16);

-- INSERT INTO ranks(name) VALUES ("rank1");
-- INSERT INTO ranks(name) VALUES ("rank2");
-- INSERT INTO ranks(name) VALUES ("rank3");
-- INSERT INTO ranks(name) VALUES ("rank4");
-- INSERT INTO ranks(name) VALUES ("rank5");
-- INSERT INTO ranks(name) VALUES ("rank6");
--
-- INSERT INTO status(name) VALUES ("status1");
-- INSERT INTO status(name) VALUES ("status2");
-- INSERT INTO status(name) VALUES ("status3");
-- INSERT INTO status(name) VALUES ("status4");
-- INSERT INTO status(name) VALUES ("status5");

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Степул Артем Мартиросович", "1997-06-12", 1, "/photo/stepul.jpg", "Поет, учавствует в студ. клубе",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "artem.stepul@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Луценко Артем Геннадійович", "1997-03-16", 1, "/photo/default.png",
        "Заместитель главы студенчиского совета факультета",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "artem.lucenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Адміністратор", "1989-09-17", 1, "/photo/admin.png", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "admin@admin.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Єнов Богдан Олександрович", "1997-06-12", 1, "/photo/enov.jpg", "Поет, учавствует в студ. клубе",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st3@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Вдовиченко Максим Ігорович", "1997-03-16", 1, "/photo/vdovichenko.jpg",
        "Заместитель главы студенчиского совета факультета",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "vdovichenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Ольшевська Ольга Володимирівна", "1989-09-17", 0, "/photo/olshevska.jpg", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380667634118", "olshevska.olga@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Іоргачов Євген Юрійович", "1998-03-01", 1, "/photo/iorgachev.jpg", "Поет, учавствует в студ. клубе",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "iorgachev@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Ткачук Станіслав Віталійович", "1997-03-16", 1, "/photo/tkachuck.jpg",
        "Заместитель главы студенчиского совета факультета",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st6@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Бодюл Олена Станіславівна", "1977-07-23", 0, "/photo/bodul.jpeg", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380487209144", "olbodiul@ukr.net",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Юрченко Боголеп Эразмович", "1970-01-01", "1", "/photo/default.png", "Старый студент",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "нету у него телефона", "st15@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        7, 0);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Клокачева Лора Протасовна", "1970-01-01", "0", "/photo/default.png", "Старый студент",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "нету у него телефона", "st16@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        7, 0);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Смирнова Катерина Василівна", "1983-12-25", 0, "/photo/smirnova.jpeg", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 14", "smirnova.kathrin@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        6, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Попков Денис Миколайович", "1982-06-03", 1, "/photo/popkov.jpg", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 14", "popkovdn@ukr.net",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        8, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Становська Тетяна Павлівна", "1951-11-25", 0, "/photo/stanovska.jpeg", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 18", "vmplotnik@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        9, 0);


INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (3, "2007-07-22", 21, 5);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (6, "2007-07-22", 40, 5);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (9, "2007-07-22", 40, 3);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (12, "2007-07-22", 40, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (13, "2007-07-22", 40, 3);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (14, "2007-07-22", 40, 5);

INSERT INTO groups(id_employee, name, id_direction)
VALUES (3, "311а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "311б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "312а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "312б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "321а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "321б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "322а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "322б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "331а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "331б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "332а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "332б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "333а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "333б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "334а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "334б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "341а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "341б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (3, "342а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "342б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "343а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (6, "343б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "344а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "344б", 1);

INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (1, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (2, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (4, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (5, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (7, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (8, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (10, "2014-09-01", 1, 1, 4, NULL);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (11, "2014-09-01", 1, 1, 4, NULL);

INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (1, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (2, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (3, 1, 0, 0, 0, 0, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (4, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (5, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (6, 0, 0, 0, 1, 0, 0, 0, 1, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (7, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (8, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (9, 0, 0, 0, 1, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (10, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (11, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (12, 0, 0, 1, 1, 1, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (13, 0, 0, 1, 1, 1, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (14, 0, 0, 0, 1, 0, 1, 0, 0, 0);


# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (1, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (2, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (3, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (4, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (5, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (6, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (7, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (8, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (9, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (10, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (11, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (12, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (13, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (14, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (15, "фыафыафыа", "122fаыфафыаasf2", "ул. фыафыа, дом Колатушафыафыакина, кварафыафыатира 12", "ыфафааыфафыаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (16, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (17, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");

INSERT INTO discipline(name)
VALUES ("ММДО");
INSERT INTO discipline(name)
VALUES ("ОАПСОС");
INSERT INTO discipline(name)
VALUES ("КПП");
INSERT INTO discipline(name)
VALUES ("ССУБД");
INSERT INTO discipline(name)
VALUES ("САПКИС");
INSERT INTO discipline(name)
VALUES ("ТКП");

INSERT INTO loads_semester(start, end, name)
VALUES (1546293600, 1562446800, "1 семестр 2019-2020 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1562446800, 1562446800, "2 семестр 2019-2020 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1562446800, 1562446800, "1 семестр 2018-2019 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1562446800, 1562446800, "2 семестр 2018-2019 року");

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (1, 9, 22, 4, 1, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (2, 9, 22, 4, 2, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (3, 12, 22, 4, 3, NULL);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (4, 13, 22, 3, 1, NULL);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (5, 14, 22, 1, 2, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (5, 14, 22, 2, 2, 9);

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 4, 13, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 4, 13, 75, "задовільно", 0, 3, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 4, 13, 75, "задовільно", 0, 2, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 2, 9, 86, "добре", 1, 2, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 2, 9, 64, "задовільно", 0, 1, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 1, 9, 65, "задовільно", 0, 1, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 4, 13, 75, "задовільно", 0, 2, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 5, 14, 86, "добре", 1, 2, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 5, 14, 64, "задовільно", 0, 1, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 1, 9, 59, "незадовільно", 1, 4, "2018-09-05");
