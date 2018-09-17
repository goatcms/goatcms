import * as utils from '../../core/utils';
import * as helpers from './helpers';

export function $renderEditBox(wrap, entity, events) {
  let form, textarea, submit;
  utils.$clean(wrap);
  helpers.$id(wrap, entity.id);
  form = document.createElement('form');
  textarea = document.createElement('textarea');
  textarea.value = entity.content;
  form.appendChild(textarea);
  submit = document.createElement('input');
  submit.type = 'submit';
  submit.value = 'save';
  form.appendChild(submit);
  wrap.appendChild(form);
  form.addEventListener('submit', events.submit);
}

export function $spin(wrap) {
  utils.$clean(wrap);
  wrap = "...";
}
