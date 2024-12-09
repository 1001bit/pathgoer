"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
class PathDeletor {
    constructor(pathId) {
        this.pathId = pathId;
        this.askedIfSure = false;
        this.initEvents();
    }
    initEvents() {
        pathDeleteButton.addEventListener("click", () => {
            if (!this.askedIfSure) {
                pathDeleteButton.innerText = "sure?";
                this.askedIfSure = true;
                return;
            }
            this.deletePath().then((err) => {
                if (err != "") {
                    pathDeleteButton.innerText = err;
                    setBorderColor(pathDeleteButton, "err");
                    return;
                }
                window.location.replace("/user");
            });
        });
        editButton.addEventListener("click", () => {
            this.cancelDelete();
        });
        pathSaveButton.addEventListener("click", () => {
            this.cancelDelete();
        });
    }
    cancelDelete() {
        this.askedIfSure = false;
        pathDeleteButton.innerText = "delete";
    }
    deletePath() {
        return __awaiter(this, void 0, void 0, function* () {
            return fetch(`/api/path/${this.pathId}`, {
                method: "DELETE",
            }).then((res) => {
                switch (res.status) {
                    case 200:
                        return "";
                    case 401:
                        return refresh().then((authd) => {
                            if (authd) {
                                return this.deletePath();
                            }
                            return "unauthorized";
                        });
                    default:
                        return "error";
                }
            });
        });
    }
}
const editSectionElem = document.getElementById("edit-section");
const pathDataElem = document.getElementById("path-data");
const editButton = document.getElementById("edit");
const pathNameInput = document.getElementById("path-name");
const pathPublicButton = document.getElementById("path-public");
const pathSaveButton = document.getElementById("save");
const pathDeleteButton = document.getElementById("delete");
const statsElem = document.getElementById("stats");
const sampleStatElem = document.getElementById("sample-stat");
const sampleEditStatElem = document.getElementById("sample-edit-stat");
class PathEditor {
    constructor(path) {
        this.path = path;
        this.initEvents();
        removeBorderColorOnFocus(pathNameInput);
    }
    initEvents() {
        editButton.addEventListener("click", () => {
            pathNameInput.value = this.path.name;
            pathPublicButton.innerText = this.path.isPublic ? "true" : "false";
            this.showSaveButtonIfChanged();
        });
        pathPublicButton.addEventListener("click", () => {
            const newPublic = pathPublicButton.innerText != "true";
            pathPublicButton.innerText = newPublic ? "true" : "false";
            this.showSaveButtonIfChanged();
        });
        pathNameInput.addEventListener("input", () => {
            this.showSaveButtonIfChanged();
        });
        pathSaveButton.addEventListener("click", () => {
            this.save();
        });
    }
    showSaveButtonIfChanged() {
        pathSaveButton.innerText = "save";
        const newName = pathNameInput.value;
        const newPublic = pathPublicButton.innerText == "true";
        if (newName == this.path.name && newPublic == this.path.isPublic) {
            setVisibility(pathSaveButton, false);
        }
        else {
            setVisibility(pathSaveButton, true);
        }
    }
    save() {
        if (pathNameInput.value == "") {
            setBorderColor(pathNameInput, "err");
            pathSaveButton.innerText = "empty name";
            return;
        }
        this.postNewData().then((err) => {
            if (err != "") {
                pathSaveButton.innerText = err;
                setBorderColor(pathNameInput, "err");
                return;
            }
            this.showSaveButtonIfChanged();
        });
    }
    postNewData() {
        return __awaiter(this, void 0, void 0, function* () {
            const newName = pathNameInput.value;
            const newPublic = pathPublicButton.innerText == "true";
            return fetch(`/api/path/${this.path.id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name: newName,
                    public: newPublic,
                }),
            }).then((res) => {
                switch (res.status) {
                    case 200:
                        this.path.name = newName;
                        this.path.isPublic = newPublic;
                        setPageTitle(newName);
                        return "";
                    case 400:
                        return "special characters";
                    case 409:
                        return "path already exists";
                    case 401:
                        return refresh().then((authd) => {
                            if (authd) {
                                return this.postNewData();
                            }
                            return "unauthorized";
                        });
                    default:
                        return "error";
                }
            });
        });
    }
}
class StatsManager {
    constructor(pathId) {
        this.pathId = pathId;
        this.pageStats = [];
    }
    renderPageStat(pageStat) {
        if (pageStat.statElem) {
            statsElem.appendChild(pageStat.statElem);
        }
        if (pageStat.editStatElem) {
            statsElem.appendChild(pageStat.editStatElem);
        }
    }
    renderStats() {
        for (const pageStat of this.pageStats) {
            this.renderPageStat(pageStat);
        }
    }
    initStats(stats, editRight) {
        if (!stats) {
            return;
        }
        for (const stat of stats) {
            const pageStat = new Stat(stat, editRight, this.pathId);
            pageStat.deletor.setDeleteCallback(() => {
                this.pageStats.splice(this.pageStats.indexOf(pageStat), 1);
            });
            this.pageStats.push(pageStat);
        }
        this.renderStats();
    }
    setEditMode(mode) {
        for (const stat of this.pageStats) {
            setVisibility(stat.statElem, !mode);
            setVisibility(stat.editStatElem, mode);
        }
    }
}
class Path {
    constructor(id) {
        this.id = id;
        this.name = "";
        this.isPublic = false;
        this.statsManager = new StatsManager(id);
        this.fetchPath();
    }
    handleFetchedData(data) {
        setPageTitle(data.path.name);
        this.statsManager.initStats(data.path.stats, data.editRight);
        if (data.editRight) {
            setVisibility(editButton, true);
            this.isPublic = data.path.public;
            this.name = data.path.name;
        }
    }
    fetchPath() {
        checkAuthAndRefresh().then((_res) => {
            fetch(`/api/path/${this.id}`, {
                method: "GET",
            }).then((res) => {
                switch (res.status) {
                    case 200:
                        res.json().then((res) => {
                            this.handleFetchedData(res);
                        });
                        break;
                    case 404:
                        window.location.replace("/404");
                        break;
                    default:
                        break;
                }
            });
        });
    }
}
const path = new Path(window.location.pathname.split("/").pop() || "0");
const deletor = new PathDeletor(path.id);
const editor = new PathEditor(path);
let editing = false;
editButton.addEventListener("click", () => {
    editing ? cancelEdit() : startEdit();
    path.statsManager.setEditMode(editing);
});
function startEdit() {
    editing = true;
    editButton.innerText = "cancel";
    setVisibility(pathDataElem, true);
}
function cancelEdit() {
    editing = false;
    editButton.innerText = "edit";
    setVisibility(pathDataElem, false);
}
class StatDeletor {
    constructor(pathId, name) {
        this.deleteCallback = () => { };
        this.pathId = pathId;
        this.name = name;
        this.askedIfSure = false;
    }
    setDeleteCallback(callback) {
        this.deleteCallback = callback;
    }
    delete() {
        if (!this.askedIfSure) {
            this.askedIfSure = true;
            return Promise.resolve("sure?");
        }
        return this.postDelete().then((message) => {
            if (message == "") {
                this.deleteCallback();
            }
            return message;
        });
    }
    postDelete() {
        return fetch(`/api/path/${this.pathId}/stat/${this.name}`, {
            method: "DELETE",
        }).then((res) => {
            switch (res.status) {
                case 200:
                    return "";
                case 401:
                    return refresh().then((authd) => {
                        if (authd) {
                            return this.postDelete();
                        }
                        return "unauthorized";
                    });
                default:
                    return "error";
            }
        });
    }
}
class StatUpdater {
    constructor(name, pathId) {
        this.name = name;
        this.pathId = pathId;
    }
    save(newName, newStepEq) {
        if (newName == "") {
            return Promise.resolve("no name");
        }
        if (Number.isNaN(newStepEq)) {
            return Promise.resolve("no step eq.");
        }
        return this.postSave({
            name: newName,
            stepEquivalent: newStepEq,
        }).then((message) => {
            if (message == "") {
                this.name = newName;
            }
            return message;
        });
    }
    postSave(newStat) {
        return fetch(`/api/path/${this.pathId}/stat/${this.name}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newStat),
        }).then((res) => {
            switch (res.status) {
                case 200:
                    return "";
                case 400:
                    return "special characters";
                case 409:
                    return "already exists";
                case 401:
                    return refresh().then((authd) => {
                        if (authd) {
                            return this.postSave(newStat);
                        }
                        return "unauthorized";
                    });
                default:
                    return "error";
            }
        });
    }
}
class Stat {
    constructor(stat, editRight, pathId) {
        this.stat = stat;
        this.statElem = this.newStatElem(this.stat);
        this.editStatElem = editRight ? this.newEditStatElem(stat) : null;
        this.deletor = new StatDeletor(pathId, stat.name);
        this.updater = new StatUpdater(stat.name, pathId);
    }
    newStatElem(stat) {
        const statElem = sampleStatElem.cloneNode(true);
        statElem.removeAttribute("id");
        setVisibility(statElem, true);
        const statNameElem = statElem.getElementsByClassName("stat-name")[0];
        const statStepEqElem = statElem.getElementsByClassName("stat-stepeq")[0];
        const countInput = new NumberInput(statElem.getElementsByClassName("stat-count-input")[0]);
        statNameElem.innerText = stat.name;
        statStepEqElem.innerText =
            "= " + stat.stepEquivalent.toString() + " steps";
        countInput.setValue(stat.count);
        return statElem;
    }
    newEditStatElem(stat) {
        const editStatElem = sampleEditStatElem.cloneNode(true);
        editStatElem.removeAttribute("id");
        const deleteButton = editStatElem.getElementsByClassName("delete-stat-button")[0];
        const saveButton = editStatElem.getElementsByClassName("save-stat-button")[0];
        const nameInput = editStatElem.getElementsByClassName("stat-name-input")[0];
        const stepEqInput = new NumberInput(editStatElem.getElementsByClassName("stat-stepeq-input")[0]);
        nameInput.value = stat.name;
        stepEqInput.setValue(stat.stepEquivalent);
        const elems = {
            deleteButton,
            saveButton,
            nameInput,
            stepEqInput,
        };
        this.initEditEvents(elems);
        return editStatElem;
    }
    initEditEvents(elems) {
        removeBorderColorOnFocus(elems.nameInput);
        removeBorderColorOnFocus(elems.stepEqInput.getInputElem());
        editButton.addEventListener("click", () => {
            elems.nameInput.value = this.stat.name;
            elems.stepEqInput.setValue(this.stat.stepEquivalent);
            this.showSaveButtonIfChanged(elems);
        });
        elems.deleteButton.addEventListener("click", () => {
            this.deletor.delete().then((message) => {
                var _a;
                if (message == "") {
                    this.statElem.remove();
                    (_a = this.editStatElem) === null || _a === void 0 ? void 0 : _a.remove();
                    return;
                }
                elems.deleteButton.innerText = message;
            });
        });
        elems.saveButton.addEventListener("click", () => {
            this.updater
                .save(elems.nameInput.value, elems.stepEqInput.getValue())
                .then((message) => {
                if (message == "") {
                    this.stat.name = elems.nameInput.value;
                    this.stat.stepEquivalent = elems.stepEqInput.getValue();
                    this.updateStat(this.stat);
                    this.showSaveButtonIfChanged(elems);
                }
            });
        });
        elems.nameInput.addEventListener("input", () => {
            this.showSaveButtonIfChanged(elems);
        });
        elems.stepEqInput.addInputListener(() => {
            this.showSaveButtonIfChanged(elems);
            removeBorderColor(elems.stepEqInput.getInputElem());
        });
    }
    updateStat(newStat) {
        this.stat = newStat;
        const statNameElem = this.statElem.getElementsByClassName("stat-name")[0];
        const statStepEqElem = this.statElem.getElementsByClassName("stat-stepeq")[0];
        statNameElem.innerText = this.stat.name;
        statStepEqElem.innerText = `= ${this.stat.stepEquivalent.toString()} steps`;
    }
    showSaveButtonIfChanged(elems) {
        const changed = !(elems.nameInput.value == this.stat.name &&
            elems.stepEqInput.getValue() == this.stat.stepEquivalent);
        elems.saveButton.innerText = "save";
        elems.deleteButton.innerText = "delete";
        setVisibility(elems.saveButton, changed);
        setVisibility(elems.deleteButton, !changed);
    }
}
function setBorderColor(elem, colorVar) {
    if (!elem) {
        return;
    }
    elem.style.borderColor = `var(--${colorVar})`;
}
function removeBorderColorOnFocus(elem) {
    if (!elem) {
        return;
    }
    elem.addEventListener("focus", () => {
        elem.style.borderColor = "";
    });
}
function removeBorderColor(elem) {
    if (!elem) {
        return;
    }
    elem.style.borderColor = "";
}
class NumberInput {
    constructor(elem) {
        this.inputElem = elem.getElementsByTagName("input")[0];
        this.plus = elem.getElementsByClassName("plus")[0];
        this.minus = elem.getElementsByClassName("minus")[0];
        this.callback = (_num) => { };
        this.initEvents();
    }
    initEvents() {
        this.plus.addEventListener("click", () => {
            this.inputElem.value = (Number(this.inputElem.value) + 1).toString();
            this.callback(Number(this.inputElem.value));
        });
        this.minus.addEventListener("click", () => {
            this.inputElem.value = (Number(this.inputElem.value) - 1).toString();
            this.callback(Number(this.inputElem.value));
        });
        this.inputElem.addEventListener("input", () => {
            this.inputElem.value = this.inputElem.value.replace(/[^0-9-]/g, "");
            this.callback(Number(this.inputElem.value));
        });
    }
    setValue(value) {
        this.inputElem.value = value.toString();
    }
    getValue() {
        return Number(this.inputElem.value);
    }
    getInputElem() {
        return this.inputElem;
    }
    addInputListener(callback) {
        this.callback = callback;
    }
}
function checkAuthAndRefresh() {
    return __awaiter(this, void 0, void 0, function* () {
        return fetch("/authenticated", {
            method: "GET",
        }).then((res) => {
            if (res.status == 401) {
                return fetch("/auth/refresh", {
                    method: "GET",
                }).then((res) => {
                    if (res.status == 200) {
                        return true;
                    }
                    return false;
                });
            }
            return true;
        });
    });
}
function refresh() {
    return __awaiter(this, void 0, void 0, function* () {
        return fetch("/auth/refresh", {
            method: "GET",
        }).then((res) => {
            if (res.status == 200) {
                return true;
            }
            return false;
        });
    });
}
const titleElem = document.getElementById("title");
function setPageTitle(title) {
    titleElem.innerText = title;
    document.title = title;
}
function setVisibility(elem, visible) {
    if (elem) {
        elem.classList.toggle("hidden", !visible);
    }
}
