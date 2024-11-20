function toggleClass(id: string, className: string) {
  const el = document.getElementById(id);

  if (el instanceof HTMLElement) {
    el.classList.toggle(className);
  }
}
