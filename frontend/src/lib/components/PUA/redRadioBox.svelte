<script lang="ts">
	export let name: string = '';
	export let id: string = '';
	export let checked: boolean = false;
</script>

<div class="flex">
	<div class="reddiobo">
		<input type="radio" {id} {name} {checked} />
		<label for={id}></label>
	</div>
</div>

<style>
	input[type='radio'] + label::before {
		content: '\a0';
		display: inline-block;
		vertical-align: middle;
		font-size: 18px;
		width: 18px;
		height: 18px;
		margin-bottom: 0.1em;
		border-radius: 50%;
		border: 2px solid #822e2e;
		text-indent: 0.15em;
		line-height: 1;
	}
	input[type='radio']:checked + label::before {
		background-color: #822e2e;
		background-clip: content-box;
		padding: 0.2em;
	}
	input[type='radio'] {
		position: absolute;
		clip: rect(0, 0, 0, 0);
	}
</style>
