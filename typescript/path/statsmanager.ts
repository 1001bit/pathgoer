/// <reference path="stat.ts" />
/// <reference path="statcreator.ts" />
/// <reference path="statcontrollers.ts" />

interface StatCount {
	name: string;
	count: number;
}

class StatsManager {
	pathId: string;
	pageStats: Array<Stat>;
	statCreator: StatCreator;

	countUpdateTicker: number;

	constructor(pathId: string) {
		this.pathId = pathId;
		this.pageStats = [];
		this.statCreator = new StatCreator(this.pathId);

		this.countUpdateTicker = 0;

		this.initEvents();
	}

	initEvents() {
		this.countUpdateTicker = setInterval(() => {
			this.updateCounts();
		}, 3000);

		window.addEventListener("beforeunload", () => {
			clearInterval(this.countUpdateTicker);
			this.updateCounts();
		});

		this.statCreator.setCreateCallback((name) => {
			const stat = {
				name: name,
				stepEquivalent: 1,
				count: 0,
			};

			const pageStat = this.initStat(stat, true);
			setVisibility(pageStat.statElem, false);
			setVisibility(pageStat.editStatElem, true);
		});
	}

	renderPageStat(pageStat: Stat) {
		statsElem.insertBefore(pageStat.statElem, statCreateBoxElem);

		if (pageStat.editStatElem) {
			statsElem.insertBefore(pageStat.editStatElem, statCreateBoxElem);
		}
	}

	initStat(stat: StatInterface, editRight: boolean): Stat {
		const pageStat = new Stat(stat, editRight, this.pathId);
		pageStat.deletor.setDeleteCallback(() => {
			this.pageStats.splice(this.pageStats.indexOf(pageStat), 1);
			this.updatePathSteps();
		});

		this.pageStats.push(pageStat);
		this.renderPageStat(pageStat);

		pageStat.stepsUpdateCallback = () => {
			this.updatePathSteps();
		};

		return pageStat;
	}

	updateCounts() {
		let counts = [];

		for (const stat of this.pageStats) {
			if (stat.newCount == stat.stat.count) {
				continue;
			}

			counts.push({
				name: stat.stat.name,
				count: stat.newCount,
			});
		}

		if (counts.length == 0) {
			return;
		}

		this.postCounts(counts);
	}

	updatePathSteps() {
		let steps = 0;
		for (const stat of this.pageStats) {
			steps += stat.newCount * stat.stat.stepEquivalent;
		}
		pathStepsElem.innerText = `steps: ${steps}`;
	}

	postCounts(counts: StatCount[]) {
		fetch(`/api/path/${this.pathId}/stats/counts`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(counts),
			keepalive: true,
		}).then((res) => {
			switch (res.status) {
				case 200:
					break;
				case 401:
					refresh().then((authd) => {
						if (authd) {
							this.postCounts(counts);
						}
					});
					break;
				default:
					console.error("Failed to update counts");
					break;
			}
		});
	}

	setEditMode(mode: boolean) {
		setVisibility(statCreateBoxElem, mode);

		for (const stat of this.pageStats) {
			setVisibility(stat.statElem, !mode);
			setVisibility(stat.editStatElem, mode);
		}
	}
}