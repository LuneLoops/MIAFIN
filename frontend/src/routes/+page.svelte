<script lang="ts">
	let count = $state(0);

	function increment() {
		count += 1;
	}
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<h1>Welcome to SvelteKit</h1>
	
	<div class="counter">
		<button onclick={increment}>
			Clicks: {count}
		</button>
	</div>

	<p>
		This is a demo SvelteKit application.
	</p>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.counter {
		display: flex;
		justify-content: center;
		margin: 2rem 0;
	}

	button {
		font-size: 1.5rem;
		padding: 1rem 2rem;
		cursor: pointer;
	}
</style>