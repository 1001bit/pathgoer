const mainElem = document.getElementsByTagName("main")[0] as HTMLDivElement;
const pathsElem = document.getElementById("paths") as HTMLDivElement;

const userStepsElem = document.getElementById("user-steps") as HTMLDivElement;
let userSteps = 0;

const userId = mainElem.getAttribute("data-user-id");

interface Stat {
	name: string;
	count: number;
	stepEquivalent: number;
}

interface Path {
	id: string;
	name: string;
	stats: Stat[];
}

function countSteps(stats: Stat[]) {
	let count = 0;

	if (stats) {
		for (const stat of stats) {
			count += stat.count * stat.stepEquivalent;
		}
	}

	userSteps += count;
	userStepsElem.innerText = `${userSteps} steps`;

	return count;
}

function newPathElem(Path: Path) {
	const samplePathElem = document.getElementById(
		"sample-path"
	) as HTMLDivElement;

	const pathElem = samplePathElem.cloneNode(true) as HTMLDivElement;
	pathElem.removeAttribute("id");
	pathElem.removeAttribute("style");

	const pathNameElem = pathElem.getElementsByClassName(
		"path-name"
	)[0] as HTMLDivElement;

	const pathLinkElem = pathElem.getElementsByClassName(
		"path-link"
	)[0] as HTMLAnchorElement;

	const pathStepsElem = pathElem.getElementsByClassName(
		"path-steps"
	)[0] as HTMLDivElement;

	pathNameElem.innerText = Path.name;
	pathLinkElem.href = `/path/${Path.id}`;
	pathStepsElem.innerText = `${countSteps(Path.stats)} steps`;
	return pathElem;
}

function renderPaths(paths: Path[]) {
	for (const path of paths) {
		const pathElem = newPathElem(path);
		// append child in front of other elements
		pathsElem.insertBefore(pathElem, pathsElem.firstChild);
	}
}

fetch(`/api/path/user/${userId}`, {
	method: "GET",
})
	.then((res) => {
		if (res.status == 200) {
			return res.json();
		} else {
			return [];
		}
	})
	.then((data) => {
		renderPaths(data);
	});