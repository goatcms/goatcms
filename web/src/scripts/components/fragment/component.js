import * as model from './model.js';
import * as utils from '../../core/utils';
import * as helpers from './helpers';
import * as view from './view';

export function init(el) {
  el.addEventListener('click', clickDispatcher);
}

function clickDispatcher(e) {
  let wrap = utils.$wrap(e.target, 'g-fragment-id');
  if(wrap===null) return;
  if(!helpers.$isOpen(wrap)) {
    $initEditBox(wrap);
    helpers.$open(wrap, true);
  }
  return null;
}

function $initEditBox(wrap) {
  let entity = helpers.$wrapEntity(wrap);
  return (entity.id === 0) ? $initNewFragmentBox(wrap, entity) : $initEditFragmentBox(wrap, entity);
}

function $initNewFragmentBox(wrap, entity) {
  entity.content = wrap.innerText;
  view.$renderEditBox(wrap, entity, {
    submit: onSubmit.bind(undefined, wrap, entity)
  });
  return Promise.resolve({});
}

function $initEditFragmentBox(wrap, entity) {
  view.$spin(wrap);
  return model.findByID(entity.id).then(resp => {
    return resp.json();
  }).then(serverEntity => {
    view.$renderEditBox(wrap, serverEntity, {
      submit: onSubmit.bind(undefined, wrap, entity)
    });
    return serverEntity;
  });
}

function onSubmit(wrap, entity, e) {
  e.preventDefault();
  entity.content = helpers.$textareaValue(wrap);
  model.persist(entity).then(resp => {
    return resp.json();
  }).then(data => {
    let id;
    if(data.id) {
      id = data.id;
      helpers.$id(wrap, id);
    } else {
      id = helpers.$id(wrap);
    }
    return utils.$serverQuerySelector('[g-fragment-id="' + id + '"]');
  }).then(serverElement => {
    utils.$clean(wrap);
    utils.$appendElementChilds(serverElement, wrap);
    helpers.$open(wrap, false);
  });
  return false;
}
