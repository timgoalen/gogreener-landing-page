package main

var GlobalCSS string = `
	* {
		margin: 0;
		padding: 0;
	}

	a {
		text-decoration: none;
		color: inherit;
	}

	body {
		height: 100vh;
		height: 100dvh;
		background: linear-gradient(0deg, #000 0%, #2E2E2E 100%);
		font-family: "Inter";
		text-align: center;
	}

	.headline {
		font-size: 1.5rem;
		font-family: "Roboto Mono";
		font-weight: 400;
		padding: 2rem;
		color: #FFF;
	}

	.get-started-container {
		display: grid;
		place-items: center;
		max-width: 94%;
		margin: auto;
	}

	.get-started {
		font-size: 1.1rem;
		font-family: "Roboto Mono";
		border: 1px solid #545454;
		border-radius: 8px;
		padding: 0.5rem;
		color: #D2D2D2;
	}

	.circle-container,
	.circle-glow-container,
	.title-container {
		position: absolute;
		top: 0;
		left: 0;
		display: grid;
		place-items: center;
		height: 100vh;
		height: 100dvh;
		width: 100%;
	}

	.circle {
		height: 400px;
		width: 400px;
		border-radius: 50%;
		display: grid;
		place-items: center;
		background: radial-gradient(circle, #000 42%, rgb(61 61 61 / 85%) 100%);
	}

	.circle-glow {
		height: 550px;
		width: 550px;
		border-radius: 50%;
		background-color: #6df36a;
		opacity: 4%;
		filter: blur(91px);
		position: absolute;
		top: 8%;
		left: 27%;
		z-index: -1;
	}

	.circle-blob {
		height: 130px;
		width: 130px;
		border-radius: 50%;
		background-color: #6df36a;
		opacity: 24%;
		filter: blur(72px);
		position: absolute;
		top: 49%;
		left: 51%;
	}

	.title {
		font-size: 3rem;
		cursor: default;
		background: linear-gradient(to right, rgba(248, 119, 0, 0.65), #FFF, rgb(55, 153, 107, 0.75));
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
	}

	.links {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 100%;
		display: flex;
		justify-content: space-evenly;
		flex-direction: row;
		padding-bottom: 2rem;
	}

	.link-container {
		display: flex;
	}

	.link-bullet-point {
		height: 24px;
		width: 24px;
		border-radius: 50%;
		background-color: #F87700;
		opacity: 22%;
		filter: blur(8px);
		margin-top: 1.1rem;
		transition: height 0.2s ease-in-out;
	}

	.link-container:hover .link-bullet-point {
		height: 48px
	}

	.link-content {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		max-width: 200px;
		margin: 1rem 0 1rem 1rem;
		cursor: pointer;
	}

	.link-title {
		font-size: 1.5rem;
		font-weight: 500;
		opacity: 85%;
		transition: opacity 0.2s ease-in-out;
		color: #FFF;
	}

	.link-body {
		text-align: left;
		font-size: 0.9rem;
		padding-top: 0.3rem;
		color: #929292;
		opacity: 85%;
		transition: opacity 0.2s ease-in-out;
	}

	.link-container:hover .link-body,
	.link-container:hover .link-title {
		opacity: 100%;
	}

	@media (max-width: 820px) {
		.links {
			flex-direction: column;
			align-items: center;
		}

		.link-content {
			max-width: 300px;
		}
	}

	@media (max-width: 400px) {
		body {
			height: 100%;
		}

		.circle-container,
		.circle-glow-container {
			display: none;
		}

		.title-container {
			position: static;
			display: grid;
			place-items: center;
			width: 100%;
			height: auto;
			margin: 2rem 0;
		}

		.links {
			position: static;
			padding-bottom: 0;
		}

		.link-container {
			max-width: 94%;
			margin: auto;
		}

		.link-bullet-point {
			display: none;
		}
	}
`
