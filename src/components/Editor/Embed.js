import axios from 'axios';

export default class Embed {
  constructor({ data, api, readOnly }) {
    this.api = api;
    this._data = {};
    this.element = null;
    this.readOnly = readOnly;

    this.data = data;
  }

  static get toolbox() {
    return {
      title: 'Embed',
      icon:
        '<svg width="13" height="14" xmlns="http://www.w3.org/2000/svg"><path d="M8.567 13.629c.728.464 1.581.65 2.41.558l-.873.873A3.722 3.722 0 1 1 4.84 9.794L6.694 7.94a3.722 3.722 0 0 1 5.256-.008L10.484 9.4a5.209 5.209 0 0 1-.017.016 1.625 1.625 0 0 0-2.29.009l-1.854 1.854a1.626 1.626 0 0 0 2.244 2.35zm2.766-7.358a3.722 3.722 0 0 0-2.41-.558l.873-.873a3.722 3.722 0 1 1 5.264 5.266l-1.854 1.854a3.722 3.722 0 0 1-5.256.008L9.416 10.5a5.2 5.2 0 0 1 .017-.016 1.625 1.625 0 0 0 2.29-.009l1.854-1.854a1.626 1.626 0 0 0-2.244-2.35z" transform="translate(-3.667 -2.7)"/></svg>',
    };
  }

  set data(data) {
    if (!(data instanceof Object)) {
      throw Error('Embed Tool data should be object');
    }

    const { html, meta } = data;

    this._data = {
      html: html || this.data.html,
      meta: meta || this.data.meta,
    };

    const oldView = this.element;

    if (oldView) {
      oldView.parentNode.replaceChild(this.render(), oldView);
    }
  }

  get data() {
    if (this.element) {
      const caption = this.element.querySelector(`.${this.api.styles.input}`);

      this._data.caption = caption ? caption.innerHTML : '';
    }

    return this._data;
  }

  render() {
    const container = document.createElement('div');
    container.setAttribute('style', 'margin: 15px');

    if (this.data.meta) {
      const { html } = this.data;

      container.innerHTML = html;
    } else {
      container.appendChild(this.makeInputHolder());
    }
    this.element = container;

    return container;
  }
  makeInputHolder() {
    const inputHolder = document.createElement('div');
    const input = document.createElement('INPUT');
    input.setAttribute('type', 'text');
    input.setAttribute('placeholder', 'Paste your link here');
    input.setAttribute('style', 'width: 100%; padding: 5px');

    // paste event requires setTimeout
    input.addEventListener('paste', (event) => {
      setTimeout(() => {
        axios
          .get('/meta', {
            params: { url: event.target.value, type: 'iframely' },
          })
          .then((res) => {
            this.data = {
              ...res.data,
            };
          })
          .catch((error) => {
            this.api.notifier.show({
              message: error.message,
              style: 'error',
            });
          });
      }, 100);
    });

    inputHolder.appendChild(input);

    return inputHolder;
  }

  save() {
    return this.data;
  }

  static get isReadOnlySupported() {
    return true;
  }

  static get sanitize() {
    return {
      html: true, // Allow HTML tags
    };
  }
}
