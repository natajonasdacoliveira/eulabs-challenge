-- +migrate Up
INSERT INTO
  product (id, name, price, description)
VALUES
  (
    null,
    'Desodorante',
    18.00,
    'Aerosol'
  ),
  (
    null,
    'Patinho de boi',
    10.00,
    'Moído'
  ),
  (
    null,
    'Tesoura de cozinha',
    10.00,
    ''
  ),
  (
    null,
    'Papel higiênico',
    10.00,
    'Folha dupla'
  ),
  (
    null,
    'Papel toalha',
    8.00,
    'Folha dupla'
  ),
  (
    null,
    'Papel alumínio',
    12.00,
    '30cm'
  ),
  (
    null,
    'Papel filme',
    15.00,
    '30cm'
  ),
  (
    null,
    'Papel manteiga',
    10.00,
    '30cm'
  ),
  (
    null,
    'Papel celofane',
    10.00,
    '30cm'
  );

-- +migrate Down
DELETE FROM
  product
WHERE
  name IN (
    'Desodorante',
    'Patinho de boi',
    'Tesoura de cozinha',
    'Papel higiênico',
    'Papel toalha',
    'Papel alumínio',
    'Papel filme',
    'Papel manteiga',
    'Papel celofane'
  );