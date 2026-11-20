-- 0002_catalog_seed.sql — curated initial catalog and safe feature defaults.
-- Idempotent by natural key (slug / key), safe to run on every boot.

INSERT OR IGNORE INTO categories (id, slug, name, description, active, sort_order, created_at, updated_at) VALUES
('cat_eletricista', 'eletricista', 'Eletricista', 'Instalação e reparos elétricos residenciais.', 1, 10, '2026-11-20T00:00:00Z', '2026-11-20T00:00:00Z'),
('cat_encanador', 'encanador', 'Encanador', 'Hidráulica, vazamentos e reparos de encanamento.', 1, 20, '2026-11-20T00:00:00Z', '2026-11-20T00:00:00Z'),
('cat_diarista', 'diarista', 'Diarista', 'Limpeza e organização residencial.', 1, 30, '2026-11-20T00:00:00Z', '2026-11-20T00:00:00Z'),
('cat_professor_particular', 'professor-particular', 'Professor particular', 'Aulas e reforço escolar presencial ou remoto.', 1, 40, '2026-11-20T00:00:00Z', '2026-11-20T00:00:00Z'),
('cat_consultor', 'consultor', 'Consultor', 'Consultoria e orientação profissional.', 1, 50, '2026-11-20T00:00:00Z', '2026-11-20T00:00:00Z');

INSERT OR IGNORE INTO category_synonyms (id, category_id, term, normalized_term, created_at) VALUES
('syn_eletricista_1', 'cat_eletricista', 'instalação elétrica', 'instalacao eletrica', '2026-11-20T00:00:00Z'),
('syn_eletricista_2', 'cat_eletricista', 'tomada', 'tomada', '2026-11-20T00:00:00Z'),
('syn_eletricista_3', 'cat_eletricista', 'fiação', 'fiacao', '2026-11-20T00:00:00Z'),
('syn_encanador_1', 'cat_encanador', 'hidráulica', 'hidraulica', '2026-11-20T00:00:00Z'),
('syn_encanador_2', 'cat_encanador', 'vazamento', 'vazamento', '2026-11-20T00:00:00Z'),
('syn_diarista_1', 'cat_diarista', 'limpeza', 'limpeza', '2026-11-20T00:00:00Z'),
('syn_diarista_2', 'cat_diarista', 'faxina', 'faxina', '2026-11-20T00:00:00Z'),
('syn_professor_1', 'cat_professor_particular', 'aula particular', 'aula particular', '2026-11-20T00:00:00Z'),
('syn_professor_2', 'cat_professor_particular', 'reforço escolar', 'reforco escolar', '2026-11-20T00:00:00Z'),
('syn_consultor_1', 'cat_consultor', 'consultoria', 'consultoria', '2026-11-20T00:00:00Z');

INSERT OR IGNORE INTO feature_flags (key, enabled, description, owner, review_at, updated_at) VALUES
('foundation.enabled', 0, 'Libera cadastro e perfil da fundação.', 'Victor', '2027-02-20', '2026-11-20T00:00:00Z'),
('admin.enabled', 0, 'Libera a área administrativa.', 'Felipe', '2027-02-20', '2026-11-20T00:00:00Z'),
('privacy.export_enabled', 0, 'Libera exportação e exclusão de dados.', 'Felipe', '2027-02-20', '2026-11-20T00:00:00Z'),
('catalog.use_synonyms', 1, 'Usa sinônimos na busca do catálogo.', 'Victor', '2027-02-20', '2026-11-20T00:00:00Z'),
('auth.password_change', 0, 'Libera troca de senha autenticada.', 'Victor', '2027-02-20', '2026-11-20T00:00:00Z'),
('auth.mfa_recovery', 0, 'Libera códigos de recuperação de MFA.', 'Felipe', '2027-02-20', '2026-11-20T00:00:00Z'),
('admin.account_requests', 0, 'Libera direitos do titular no admin.', 'Felipe', '2027-02-20', '2026-11-20T00:00:00Z'),
('profile.precise_location', 0, 'Libera localização precisa cifrada.', 'Felipe', '2027-02-20', '2026-11-20T00:00:00Z');
