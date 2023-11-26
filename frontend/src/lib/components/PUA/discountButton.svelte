<script lang="ts">
	export let text: string = 'null';
</script>

<div id="radio">
	<label
		><input type="radio" name="variety" value="" />
		<span class="round button">{text}</span></label
	>
</div>

<style>
	#radio input[type='radio'] {
		display: none;
	}
	#radio input:checked + .button {
		background: #487444;
		color: white;
	}
	#radio .button {
		display: inline-block;
		margin: 0 5px 10px 0;
		padding: 5px 10px;
		background: transparent;
		cursor: pointer;
		color: #287821;
		font-weight: bold;
		padding-top: 0px;
		padding-bottom: 0px;
	}
	#radio .button:hover {
		border-color: #287821;
	}
	#radio .round {
		border-radius: 10px;
		border: 2px solid #2d9a25;
		border-color: #287821;
	}
</style>
