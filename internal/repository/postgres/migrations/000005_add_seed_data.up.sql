-- Сид данных для команд
INSERT INTO team (team_name) VALUES
                                 ('Backend'),
                                 ('Frontend'),
                                 ('DevOps');

-- Сид данных для пользователей
INSERT INTO users (id, username, team_name, is_active) VALUES
-- Backend команда
(gen_random_uuid(), 'alice', 'Backend', TRUE),
(gen_random_uuid(), 'eve', 'Backend', TRUE),
(gen_random_uuid(), 'henry', 'Backend', TRUE),
(gen_random_uuid(), 'ivy', 'Backend', TRUE),

-- Frontend команда
(gen_random_uuid(), 'bob', 'Frontend', TRUE),
(gen_random_uuid(), 'frank', 'Frontend', TRUE),
(gen_random_uuid(), 'jack', 'Frontend', TRUE),
(gen_random_uuid(), 'kate', 'Frontend', TRUE),

-- DevOps команда
(gen_random_uuid(), 'carol', 'DevOps', TRUE),
(gen_random_uuid(), 'grace', 'DevOps', TRUE),
(gen_random_uuid(), 'louis', 'DevOps', TRUE),
(gen_random_uuid(), 'mia', 'DevOps', TRUE),

-- Пользователи без команды
(gen_random_uuid(), 'dave', NULL, TRUE),
(gen_random_uuid(), 'nina', NULL, TRUE);

-- Сид данных для pull requests
INSERT INTO pull_request (id, "name", author_id, status, created_at, merged_at)
VALUES
    ('PR-001', 'Add login endpoint', (SELECT id FROM users WHERE username='alice'), 'OPEN', NOW() - INTERVAL '5 days', NULL),
    ('PR-002', 'Fix navbar bug', (SELECT id FROM users WHERE username='bob'), 'MERGED', NOW() - INTERVAL '10 days', NOW() - INTERVAL '2 days'),
    ('PR-003', 'Update CI pipeline', (SELECT id FROM users WHERE username='carol'), 'OPEN', NOW() - INTERVAL '1 days', NULL),
    ('PR-004', 'Refactor auth service', (SELECT id FROM users WHERE username='henry'), 'OPEN', NOW() - INTERVAL '3 days', NULL),
    ('PR-005', 'Improve caching', (SELECT id FROM users WHERE username='jack'), 'OPEN', NOW() - INTERVAL '2 days', NULL);

-- Сид данных для pull request reviewers
-- 2 ревьюера из команды автора PR
-- PR-001 автор alice (Backend) -> possible reviewers: eve, henry, ivy
INSERT INTO pull_request_reviewer (pull_request_id, reviewer_id, assigned_at)
VALUES
    ('PR-001', (SELECT id FROM users WHERE username='eve'), NOW() - INTERVAL '4 days'),
    ('PR-001', (SELECT id FROM users WHERE username='henry'), NOW() - INTERVAL '4 days'),

-- PR-002 автор bob (Frontend) -> possible reviewers: frank, jack, kate
    ('PR-002', (SELECT id FROM users WHERE username='frank'), NOW() - INTERVAL '9 days'),
    ('PR-002', (SELECT id FROM users WHERE username='jack'), NOW() - INTERVAL '9 days'),

-- PR-003 автор carol (DevOps) -> possible reviewers: grace, louis, mia
    ('PR-003', (SELECT id FROM users WHERE username='grace'), NOW() - INTERVAL '1 days'),
    ('PR-003', (SELECT id FROM users WHERE username='louis'), NOW() - INTERVAL '1 days'),

-- PR-004 автор henry (Backend) -> reviewers: alice, eve
    ('PR-004', (SELECT id FROM users WHERE username='alice'), NOW() - INTERVAL '2 days'),
    ('PR-004', (SELECT id FROM users WHERE username='eve'), NOW() - INTERVAL '2 days'),

-- PR-005 автор jack (Frontend) -> reviewers: bob, frank
    ('PR-005', (SELECT id FROM users WHERE username='bob'), NOW() - INTERVAL '1 days'),
    ('PR-005', (SELECT id FROM users WHERE username='frank'), NOW() - INTERVAL '1 days');
