export function $wrap(el, attr) {
  while(el !== document && el !== null) {
    if(el.getAttribute(attr)) {
      return el;
    }
    el = el.parentNode;
  }
  return null;
}

export function $clean(el) {
  while (el.firstChild) {
      el.removeChild(el.firstChild);
  }
}

export function $appendElementChilds(from, to) {
  let c = from.children;
  for (let i = 0; i < c.length; i++) {
    to.appendChild(c[i]);
  }
}

export function $serverQuerySelector(query) {
  return fetch(window.location.href).then(resp => {
    return resp.text();
  }).then(html => {
    let wrap = document.createElement('div');
    wrap.innerHTML = html;
    return wrap.querySelector(query);
  });
}

export function $lang() {
  return document.documentElement.lang;
}
