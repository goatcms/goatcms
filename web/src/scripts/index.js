import '../styles/index.scss';
import * as FragmentComponent from './components/fragment/component';

document.addEventListener("DOMContentLoaded", () => {
  let body = document.getElementsByTagName("BODY")[0];
  FragmentComponent.init(body);
});
