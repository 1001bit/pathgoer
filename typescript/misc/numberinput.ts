class NumberInput {
	private elem: HTMLDivElement;
	private inputElem: HTMLInputElement;
	private plus: HTMLDivElement;
	private minus: HTMLDivElement;
	private callback: (num: number) => void;

	constructor(elem: HTMLDivElement) {
		this.elem = elem;
		this.inputElem = elem.getElementsByTagName(
			"input"
		)[0] as HTMLInputElement;

		this.plus = elem.getElementsByClassName("plus")[0] as HTMLDivElement;
		this.minus = elem.getElementsByClassName("minus")[0] as HTMLDivElement;

		this.callback = (_num: number) => {};
		this.initEvents();
	}

	private initEvents() {
		this.plus.addEventListener("click", () => {
			this.inputElem.value = (
				Number(this.inputElem.value) + 1
			).toString();

			this.callback(Number(this.inputElem.value));
		});
		this.minus.addEventListener("click", () => {
			this.inputElem.value = (
				Number(this.inputElem.value) - 1
			).toString();

			this.callback(Number(this.inputElem.value));
		});

		this.inputElem.addEventListener("input", () => {
			this.inputElem.value = this.inputElem.value.replace(/[^0-9-]/g, "");

			this.callback(Number(this.inputElem.value));
		});
	}

	setValue(value: number) {
		this.inputElem.value = value.toString();
	}

	getValue() {
		return Number(this.inputElem.value);
	}

	getElem() {
		return this.elem;
	}

	getInputElem() {
		return this.inputElem;
	}

	addInputListener(callback: (num: number) => void) {
		this.callback = callback;
	}
}

function acceptOnlyNumbers(elem: HTMLInputElement) {
	elem.addEventListener("input", () => {
		elem.value = elem.value.replace(/[^0-9-]/g, "");
	});
}
