// ESILVA.ORG UTILS
function getActiveTabId() {
  const activeBtn = document.querySelector("#page-navigation button.selected");
  return activeBtn.id;
}

function openCollapsibleFast(collapsible) {
  collapsible.classList.add("open");
  let content = collapsible.nextElementSibling;
  content.style.maxHeight = content.scrollHeight + content.scrollHeight + "px";
}
// ESILVA.ORG UTILS

// ESILVA.ORG CACHE
class esilvaCache {
  #cachePath = "esilva-org-cache";
  #cacheObj;

  constructor() {
    this.#getLocalStorage();
    // If the localStorage already holds cache of ours we don't need to
    // initialize it
    if (this.#cacheObj) {
      return;
    }

    this.#initializeCacheObj();
    this.#updateLocalStorage();
  }

  // Initializes the cache object. I decided to have a JSON that will store
  // data keyed by each tab and then content type, for example:
  // obj = {
  //   "projects": {
  //     "collapsibles": {
  //       "collapsible1_id": value,
  //     },
  //   },
  // }
  #initializeCacheObj() {
    this.#cacheObj = new Object;
    const btns = document.querySelectorAll("#page-navigation button");
    for (let k = 0; k < btns.length; k++) {
      this.#cacheObj[btns[k].id] = new Object;
      this.#cacheObj[btns[k].id]["collapsibles"] = new Object;
    }

    const activeTab = getActiveTabId();
    const collapsibles = document.querySelectorAll(".collapsible");
    for (let c = 0; c < collapsibles.length; c++) {
      if (collapsibles[c].classList.contains("open")) {
        this.#cacheObj[activeTab]["collapsibles"][collapsibles[c].id] = true;
      } else {
        this.#cacheObj[activeTab]["collapsibles"][collapsibles[c].id] = false;
      }
    }
  }

  // Updates the localStorage cache variable with the data that this objects
  // cache variable is holding
  #updateLocalStorage() {
    localStorage.setItem(this.#cachePath, JSON.stringify(this.#cacheObj));
  }

  // Grabs the localStorage data regarding our cache and loads it into
  // this objects cache variable
  #getLocalStorage() {
    this.#cacheObj = JSON.parse(localStorage.getItem(this.#cachePath));
  }

  // This method will load the state that is stored in localStorage into the
  // website.
  loadState() {
    this.#getLocalStorage();

    // Load collapsibles states 
    for (let tab in this.#cacheObj) {
      if (!this.#cacheObj[tab]["collapsibles"]) {
        continue;
      }

      for (let collapsible in this.#cacheObj[tab]["collapsibles"]) {
        if (this.#cacheObj[tab]["collapsibles"][collapsible]) {
          let elem = document.getElementById(collapsible);
          if (elem) {
            openCollapsibleFast(elem);
          }
        }
      }
    }
  }

  // We might be able to make this more elegant, that will be a later task
  toggleCollapsibleState(cId) {
    console.log(cId);

    const tabId = getActiveTabId();
    this.#getLocalStorage();

    if (!this.#cacheObj[tabId]["collapsibles"][cId]) {
      // If the collapsible wasn't present in the object, then we assume that
      // it was closed since we have no memory of its state, therefore we have
      // to open it
      this.#cacheObj[tabId]["collapsibles"][cId] = false;
    }

    // Toggle the state of the collapsible in cache
    if (this.#cacheObj[tabId]["collapsibles"][cId] == true) {
      this.#cacheObj[tabId]["collapsibles"][cId] = false;
    } else {
      this.#cacheObj[tabId]["collapsibles"][cId] = true;
    }

    this.#updateLocalStorage();
  }
}
// Global object to work on our cache
cache = new esilvaCache;
// ESILVA.ORG CACHE


// The buttons are also dynamically loaded so they have to be delegated
function addEventNavListeners() {
  navButtons = document.querySelectorAll("button.nav-button");

  document.addEventListener("click", function(e) {
    const target = e.target.closest(".nav-button");
    if (!target) {
      return;
    }

    // De-select the current selected button
    selNavBtn = document.querySelector("button.nav-button.selected");
    selNavBtn.classList.remove("selected");

    // Select the new button
    target.classList.add("selected");
  });
}

function addImgEventListeners() {
  // This content is dynamically loaded, we have to delegate the events
  document.addEventListener("click", function(e) {
    const target = e.target.closest("img");
    if (target) {
      window.open(target.src, '_blank');
    }
  });
}

/*
  * I wish I could make this one better, but for now I think it works
  * I really want to normalize all the line heights so that I can achieve
  * true terminal scrolling.
 */
function customScroller() {
  // This function cycles trough the available tabs given the direction ammount
  // positive dir moves N times to the right
  // negative dir moves N times to the left
  function moveTab(dir) {
    btns = document.querySelectorAll("#page-navigation button");

    for (let k = 0; k < btns.length; k++) {
      if (btns[k].classList.contains("selected")) {
        var clickEvent = new MouseEvent("click", {
          "view": window,
          "bubbles": true,
          "cancelable": false
        });

        dst = k + dir;
        if (dst >= btns.length) {
          dst = 0;
        } else if (dst < 0) {
          dst = btns.length - 1;
        }

        btns[dst].dispatchEvent(clickEvent);
        break;
      }
    }
  }

  // This might give us problems if the page changes sizes
  fontHeight = parseInt(
    window.getComputedStyle(document.getElementById("content")).fontSize, 10
  );
  elem = document.getElementById("main-container-content");

  document.addEventListener("keydown", (e) => {
    const key = e.key;
    // switch (key) {
    //   case 'j':
    //     elem.scrollBy(0, fontHeight);
    //     break;
    //   case 'k':
    //     elem.scrollBy(0, -fontHeight);
    //     break;
    //   case 'g':
    //     elem.scrollTo(0, 0)
    //     break;
    //   case 'G':
    //     elem.scrollTo(0, elem.scrollHeight)
    //     break;
    //   case 'l':
    //     moveTab(1)
    //     break;
    //   case 'h':
    //     moveTab(-1)
    //     break;
    // }
  });
}


let count = 0;
let originalContent = "";
const data = ["\\", "|", "/", "―", "\\", "|", "/", "―"];
// const data = [ ".", "•", "°", "*", "" ];
// const data = ["┤", "┘", "┴", "└", "├", "┌", "┬", "┐",];
// const data = ["◜ ", " ◝", " ◞", "◟ "];
// ⣿
// const data = ["⡀", "⠄", "⠂", "⠁", "⠈", "⠐", "⠠", "⢀"];
/* const data = ["⠁", "⠂", "⠄", "⡈", "⡐", "⡠", "⣀", "⣁", "⣂", "⣄", "⣌", "⣔", "⣤",
            "⣥", "⣦", "⣮", "⣶", "⣷", "⣿", "⣶", "⣤", "⣀", " "]; */

async function spinner(elem) {
  originalContent = elem.innerHTML;
  while (elem.classList.contains("loading")) {
    elem.innerHTML = originalContent + " " + data[count];
    count += 1;
    if (count >= data.length) {
      count = 0;
    }

    await new Promise(r => setTimeout(r, 200));
  }

  count = 0;
  elem.innerHTML = originalContent
}

// 
function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

/*
Recursive function to render the nodes in a div.
*/
async function renderChildNodes(root, parentNode, childNode) {
  if (childNode.nodeName === "#text") {
    for (let c = 0; c < childNode.data.length;) {
      await new Promise(r => setTimeout(r, 1));

      let charsToRender = getRandomInt(1, 15);
      let diff = childNode.data.length - charsToRender;

      if (diff < 0) {
        charsToRender += diff;
      }

      let strToAppend = childNode.data.substring(c, c + charsToRender);
      console.log(strToAppend);

      parentNode.append(strToAppend);
      root.style.maxHeight = root.scrollHeight + parentNode.scrollHeight + "px";

      c += charsToRender;
    }
  } else {
    let newElem = document.createElement(childNode.tagName);

    newElem.attributes = childNode.attributes;
    for (let i = 0; i < childNode.attributes.length; i++) {
      let attr = childNode.attributes[i];
      newElem.setAttribute(attr.name, attr.value);
    }

    parentNode.append(newElem);
    root.style.maxHeight = root.scrollHeight + parentNode.scrollHeight + "px";

    for (let i = 0; i < childNode.childNodes.length; i++) {
      await renderChildNodes(root, newElem, childNode.childNodes[i]);
    }
  }
}

gCollapsibles = new Object;
function addCollapsiblesEventListeners() {
  // Add event delegation
  document.addEventListener("click", function(e) {
    const target = e.target.closest(".collapsible");
    // If we click on nothing it will log an error. We don't need that.
    if (!target) {
      return
    }

    if (!gCollapsibles[target.id]) {
      gCollapsibles[target.id] = new esilvaCollapsible(target.id)
    }

    gCollapsibles[target.id].click();
  });
}

class esilvaCollapsible {
  #collapsibleID;
  // #collapsibleOriginalContent;

  #collapsible;
  #content;
  #rendering;

  constructor(cId) {
    // Store the "metadata". The IDs and how to access it on the DOM
    this.#collapsibleID = cId;
    this.#rendering = false;
  }

  #renderingCleanUp() {
    this.#collapsible.classList.toggle("open");
    this.#collapsible.classList.toggle("loading");
    cache.toggleCollapsibleState(this.#collapsible.id);
    this.#rendering = false;
  }

  #close() {
    this.#content.style.maxHeight = null;

    // Cleaning up
    this.#renderingCleanUp()
  }

  async #open() {
    // We get the childnodes that are currently hidden into memory
    let origContent = Array.from(this.#content.childNodes);

    // And then set the content to nothing so that we can start filling it
    this.#content.innerHTML = ""

    // Render every child node recursively
    for (let i = 0; i < origContent.length; i++) {
      await renderChildNodes(this.#content, this.#content, origContent[i]);
    }

    this.#renderingCleanUp()
  }

  #isOpen() {
    // The element is rendered if it has a set maxHeight or if it has the class
    // open
    return this.#content.style.maxHeight || this.#content.classList.contains("open");
  }

  async click() {
    // Do not allow to execute rendering while a rendering action is already
    // under way
    if (this.#rendering) {
      return;
    }

    // Set the state to rendering:
    // - object rendering to true;
    // - collapsible goes into loading;
    // - attach the spinner;
    this.#rendering = true;

    // Get the collapsible element
    this.#collapsible = document.getElementById(this.#collapsibleID)
      .closest(".collapsible");

    // Get the collapsibles content
    this.#content = this.#collapsible.nextElementSibling;

    // Set the collapsible to loading
    this.#collapsible.classList.toggle("loading");
    spinner(this.#collapsible);

    // Perform the appropriate rendering action
    if (this.#isOpen()) {
      this.#close();
    } else {
      this.#open();
    }
  }
}

// main function
document.addEventListener("DOMContentLoaded", function() {
  cache.loadState();

  addEventNavListeners();
  addImgEventListeners();
  addCollapsiblesEventListeners();
  customScroller();
})

function navigateToPage(event) {
  // Prevent event propagation to prevent triggering collapsible event listeners
  // When we click to go to the resource
  event.stopPropagation();

  // Perform navigation to the desired page
  // Example: window.location.href = 'your_page_url';
}

