import { writable } from "svelte/store";

const varStore = {
  variables: [],
  callbacks: [],
  getVar: function (name) {
    const val = this.variables.filter((item) => item.name === name);
    if (val.length > 0 && typeof val[0].value !== "undefined")
      return val[0].value;
    return null;
  },
  setVar: function (name, varValue) {
    let oldvarInd = this.variables.findIndex((item) => item.name === name);
    if (oldvarInd === -1) {
      this.variables.push({
        name: name,
        value: varValue,
      });
    } else {
      if (varValue !== null) this.variables[oldvarInd].value = varValue;
    }
    const cb = this.callbacks.filter((item) => item.name === name);
    if (cb.length > 0 && typeof cb[0].value === "function") cb[0].value();
  },
  setUpVar: function (name, def, callback) {
    this.callbacks = this.callbacks.filter((item) => item.name !== name);
    this.setVar(name, def);
    this.callbacks.push({
      name: name,
      value: callback,
    });
  },
};

export const variables = writable(varStore);
