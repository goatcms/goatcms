wrapimport * as model from './model.js';
import * as utils from '../../core/utils';

export function init(el) {
  el.addEventListener('click', clickDispatcher);
}

function clickDispatcher(e) {
  let wrap = utils.$wrap(e.target, 'g-fragment-id');
  if(wrap===null) return;
  if($isOpen(wrap)) {
  } else {
    $initEditBox(wrap)
    $open(wrap, true);
  }
  return null;
}

function $initEditBox(wrap) {
  let form, entity, textarea, submit, id
  utils.$clean(wrap);
  form = document.createElement('form');
  textarea = document.createElement('textarea');
  textarea.readOnly = true;
  form.appendChild(textarea);
  submit = document.createElment('input');
  submit.type = 'submit';
  submit.value = 'save';
  textarea.disabled = true;
  form.appendChild(submit);
  form.addEventListener('submit', (e) => {
    e.preventDefault();
    entity = $wrapEntity(e.target);
    entity.content = textarea.value;
    model.persist(entity).then(data => {
      //TODO: update id for insert
      debugger
      if(data.id) $id(wrap, data.id);
      return utils.$serverQuerySelector('[g-fragment-id="' + data.id + '"]')
    }).then(serverElement => {
      utils.$clean(wrap);
      utils.$appendElementChilds(serverElement, wrap)
      wap.innerHTML = data.content;
    });
    return false;
  });
  id = $id(wrap);
  return model.findByID(id).then(entity => {
    textarea.value = entity.content;
    textarea.readOnly = false;
    submit.disabled = false;
    return entity;
  });
}

function $isOpen(el) {
  return el.getAttribute('g-fragment-opened') === 'y';
}

function $open(el, state) {
  el.setAttribute('g-fragment-opened', state ? 'y' : 'n')
}

function $id(el, value) {
  if(value !== undefined) {
    el.setAttribute('g-fragment-id', value)
  }
  return parseInt(el.getAttribute('g-fragment-id'));
}

function $wrapEntity(el) {
  let wrap = utils.$wrap(el, 'g-fragment-id');
  if(wrap === null) {
    return null;
  }
  let id = wrap.getAttribute('g-fragment-id');
  let key = wrap.getAttribute('g-fragment-key');
  let keysplit = key.split('.', 2);
  return {
    id: parseInt(id) || 0,
    lang: keysplit[0],
    name: keysplit[1]
  }
}
