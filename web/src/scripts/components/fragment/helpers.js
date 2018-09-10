import * as utils from '../../core/utils';

export function $isOpen(el) {
  return el.getAttribute('g-fragment-opened') === 'y';
}

export function $open(el, state) {
  el.setAttribute('g-fragment-opened', state ? 'y' : 'n');
}

export function $id(el, value) {
  if(value !== undefined) {
    el.setAttribute('g-fragment-id', value);
  }
  return parseInt(el.getAttribute('g-fragment-id'));
}

export function $wrapEntity(el) {
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
  };
}

export function $textareaValue(wrap) {
  return wrap.querySelector('textarea').value;
}
