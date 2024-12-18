/// <reference path="elems.ts" />

const avatarInput = document.createElement("input");
avatarInput.type = "file";
avatarInput.accept = "image/*";
setVisibility(avatarInput, false);

changeAvatarElem.addEventListener("click", () => {
	avatarInput.click();
});

avatarInput.addEventListener("change", async () => {
	const file = avatarInput.files?.item(0);
	if (!file) {
		return;
	}

	const formData = new FormData();
	formData.append("avatar", file);

	try {
		fetch("/dynamic/avatar", {
			method: "POST",
			body: formData,
		}).then((res) => {
			if (!res.ok) {
				changeAvatarElem.innerHTML = "error";
				return;
			}

			avatarElem.src = avatarElem.src.split("?")[0] + "?" + Date.now();
		});
	} catch (e) {
		changeAvatarElem.innerHTML = "error";
	}
});
