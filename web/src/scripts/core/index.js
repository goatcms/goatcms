export function $wrap(el, attr) {
  while(el !== document && el !== null) {
    if(el.getAttribute(attr)) {
      return el
    }
    el = el.parentNode;
  }
  return null
}

export function $clear(el) {
  while (el.firstChild) {
      el.removeChild(el.firstChild);
  }
}

export function $lang() {
  return document.documentElement.lang;
}
