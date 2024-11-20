"use strict";
function toggleClass(id, className) {
    const el = document.getElementById(id);
    if (el instanceof HTMLElement) {
        el.classList.toggle(className);
    }
}
