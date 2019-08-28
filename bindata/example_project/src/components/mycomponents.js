const req = require.context(".", false, /\.vue$/);
const components = {};
req.keys().forEach(fileName => {
  const componentName = fileName.replace(/(\.\/|\.vue)/g, "");
  components[componentName] = req(fileName).default;
});
const mycomponents = {
  install: function(Vue) {
    for (var key in components) {
      Vue.component(key, components[key]);
    }
  }
};

export default mycomponents;
