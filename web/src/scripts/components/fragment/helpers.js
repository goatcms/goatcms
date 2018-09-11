import * as utils from '../../core/utils';

export function $isOpen(el) {
  return el.classList.contains('open');
  //el.getAttribute('g-fragment-opened') === 'y';
}

export function $open(el, state) {
  if(state) {
    el.classList.add('open');
  } else {
    el.classList.remove('open');
  }
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
  let lang = key.split('.', 1)[0];
  let name = key.substr(lang.length+1);
  return {
    id: parseInt(id) || 0,
    lang: lang,
    name: name
  };
}

export function $textareaValue(wrap) {
  return wrap.querySelector('textarea').value;
}
