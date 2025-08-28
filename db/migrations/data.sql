-- Insert into houses
INSERT INTO houses (id, first_name, last_name, head_of_house) VALUES
('gryf-001', 'Gryffindor', 'House', 'Minerva McGonagall'),
('huff-001', 'Hufflepuff', 'House', 'Pomona Sprout'),
('rave-001', 'Ravenclaw', 'House', 'Filius Flitwick'),
('slyt-001', 'Slytherin', 'House', 'Severus Snape');

-- Insert into students
INSERT INTO students (id, first_name, last_name, age, house_id) VALUES
('stu-001', 'Harry', 'Potter', 11, 'gryf-001'),
('stu-002', 'Hermione', 'Granger', 11, 'gryf-001'),
('stu-003', 'Ron', 'Weasley', 11, 'gryf-001'),
('stu-004', 'Draco', 'Malfoy', 11, 'slyt-001'),
('stu-005', 'Luna', 'Lovegood', 11, 'rave-001'),
('stu-006', 'Cedric', 'Diggory', 12, 'huff-001'),
('stu-007', 'Neville', 'Longbottom', 11, 'gryf-001'),
('stu-008', 'Cho', 'Chang', 11, 'rave-001'),
('stu-009', 'Pansy', 'Parkinson', 11, 'slyt-001'),
('stu-010', 'Susan', 'Bones', 11, 'huff-001');
