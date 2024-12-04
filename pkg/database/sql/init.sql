-- adiciona organizacao admin
INSERT INTO public.organizacao
("nome", "id_assinatura") VALUES ('Organização XPTO', 'stripe-admin-id') ON CONFLICT DO NOTHING;

-- adiciona organizacao de um usuário master
INSERT INTO public.organizacao
("nome", "id_assinatura")
VALUES
('Fulanos Advogados', 'stripe-master-id') ON CONFLICT DO NOTHING;

-- cria Permissoes padrao do sistema
INSERT INTO public.permissao
(nome, ativo) VALUES ('Gerar Documentos', true);

INSERT INTO public.permissao
(nome, ativo) VALUES ('Baixar DOCX', true);

INSERT INTO public.permissao
(nome, ativo) VALUES ('Baixar PDF', true);


-- cria perfil admin
INSERT INTO public.perfil
("id_organizacao",
"criado_em",
"ativo",
"descricao",
"master",
"nome")
VALUES
(1,
current_timestamp,
TRUE,
'Descrição Perfil Admin',
TRUE,
'Admin') ON CONFLICT DO NOTHING;

-- cria perfil master
INSERT INTO public.perfil
("id_organizacao",
"criado_em",
"ativo",
"descricao",
"master",
"nome")
VALUES
(1,
current_timestamp,
TRUE,
'Descrição Perfil Master',
TRUE,
'Master') ON CONFLICT DO NOTHING;

-- cria perfil sócio
INSERT INTO public.perfil
("id_organizacao",
"criado_em",
"ativo",
"descricao",
"master",
"nome")
VALUES
(1,
current_timestamp,
TRUE,
'Descrição Perfil Sócio',
FALSE,
'Sócio') ON CONFLICT DO NOTHING;

-- cria perfil colaborador
INSERT INTO public.perfil
("id_organizacao",
"criado_em",
"ativo",
"descricao",
"master",
"nome")
VALUES
(1,
current_timestamp,
TRUE,
'Descrição Perfil Colaborador',
FALSE,
'Colaborador') ON CONFLICT DO NOTHING;

-- cria perfil estagiário
INSERT INTO public.perfil
("id_organizacao",
"criado_em",
"ativo",
"descricao",
"master",
"nome")
VALUES
(1,
current_timestamp,
TRUE,
'Descrição Perfil Estagiário',
FALSE,
'Estagiário') ON CONFLICT DO NOTHING;

--atrelar permissoes aos perfis
INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(1,
1) ON CONFLICT DO NOTHING;

INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(1,
2) ON CONFLICT DO NOTHING;

INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(1,
3) ON CONFLICT DO NOTHING;

INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(2,
1) ON CONFLICT DO NOTHING;

INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(2,
2) ON CONFLICT DO NOTHING;

INSERT INTO public.perfil_permissoes
("id_perfil",
"id_permissao")
VALUES
(2,
3) ON CONFLICT DO NOTHING;

-- cria usuario master
INSERT INTO public.usuario
("id_organizacao",
"id_provedor_autenticacao",
"criado_em",
"email",
"nome",
"telefone",
"status",
"master",
"id_perfil")
VALUES
(1,
'1uWuuuVE2jO9DFdecATwKWepjou1',
current_timestamp,
'master@gmail.com',
'nome master',
'81999999999',
1,
TRUE,
2) ON CONFLICT DO NOTHING;

-- cria usuario sócio
INSERT INTO public.usuario
("id_organizacao",
"id_provedor_autenticacao",
"criado_em",
"email",
"nome",
"telefone",
"status",
"master",
"id_perfil")
VALUES
(1,
'0MYVJi4PHFOleoCqvb2EX5crxbn2',
current_timestamp,
'socio@gmail.com',
'nome socio',
'81999999999',
1,
FALSE,
3) ON CONFLICT DO NOTHING;

-- cria usuario Colaborador
INSERT INTO public.usuario
("id_organizacao",
"id_provedor_autenticacao",
"criado_em",
"email",
"nome",
"telefone",
"status",
"master",
"id_perfil")
VALUES
(1,
'jjh8JsG90qhNECM1uoobmpHgCSJ2',
current_timestamp,
'colaborador@gmail.com',
'nome colaborador',
'81999999999',
2,
FALSE,
4) ON CONFLICT DO NOTHING;

-- cria usuario estagiário
INSERT INTO public.usuario
("id_organizacao",
"id_provedor_autenticacao",
"criado_em",
"email",
"nome",
"telefone",
"status",
"master",
"id_perfil")
VALUES
(1,
'Y7UWCAZrPaRtgsPUnEIcizUNFMl1',
current_timestamp,
'estagiario@gmail.com',
'nome estagiario',
'81999999999',
0,
FALSE,
5) ON CONFLICT DO NOTHING;
