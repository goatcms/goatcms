export function findByID(id) {
  return fetch("/rest/model/fragment/" + id);
}

export function create(entity) {
  return fetch('/rest/model/fragment', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(entity)
  }).then(data => {
    entity.id = data.id;
    return data;
  });
}

export function update(entity) {
  if(!entity.id) {
    throw "incorrect id " + entity.id;
  }
  return fetch('/rest/model/fragment/' + entity.id, {
    method: 'PUT',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(entity)
  });
}

export function persist(entity) {
  if(entity.id) {
    return update(entity);
  }
  return create({
    name: entity.name,
    lang: entity.lang,
    content: entity.content,
  });
}
