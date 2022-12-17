import { writable } from "svelte/store";

var timerObj = {
  Callbacks: [],
  TimerId: null,
  time: 0,
  AddTimer: function (name, callback) {
    //
    // Remove the former callback if it exists.
    //
    timerObj.Callbacks = timerObj.Callbacks.filter(
      (item) => item.name !== name
    );

    //
    // Add the callback.
    //
    timerObj.Callbacks.push({
      name: name,
      callback: callback,
    });
  },
  StartTimer: function () {
    //
    // Set the next timer.
    //
    //timerObj.TimerId = setTimeout(timerObj.NextTime, 60000);
  },
  NextTime: function () {
    //
    // Increment the timer.
    //
    timerObj.time = timerObj.time + 1;

    //
    // Call the callbacks with the time.
    //
    if (timerObj.Callbacks.length > 0) {
      for (let i = 0; i < timerObj.Callbacks.length; i++) {
        timerObj.Callbacks[i].callback(timerObj.time);
      }
    }

    //
    // Set the next timer.
    //
    timerObj.TimerId = setTimeout(timerObj.NextTime, 60000);
  },
  StopTimer: function () {
    if (timerObj.TimerId !== null) {
      clearTimeout(timerObj.TimerId);
      timerObj.time = 0;
    }
  },
  ClearTimers: function () {
    timerObj.Callbacks = [];
    timerObj.time = 0;
  },
};
export const timer = writable(timerObj);
