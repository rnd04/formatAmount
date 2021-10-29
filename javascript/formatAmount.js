function formatAmount(number, decimal) {
	if (number === 0 || number === '0') {
		return '0';
	}
	let sign = parseInt(number) < 0 ? '-' : '';
	let absstr = Math.abs(number).toString();
	let monetary = '0';
	let fraction;
	if (absstr.length > decimal) {
		let integer = absstr.substring(0, absstr.length - decimal);
		fraction = absstr.substring(absstr.length - decimal);
		let i1 = integer.length % 3;
		let m = i1 == 0 ? [] : [integer.substring(0, i1)];
		for (let i2 = 0; i2 < (integer.length - i1) / 3; i2++) {
			m.push(integer.substring(i1 + i2 * 3, i1 + i2 * 3 + 3));
		}
		monetary = m.join(',');
	} else if (absstr.length == decimal) {
		fraction = absstr;
	} else {
		fraction = absstr.padStart(decimal, '0');
	}
	fraction = fraction.replace(/0+$/, ''); // rtrim
	let p = fraction.length == 0 ? '' : '.'; // decimal point
	return `${sign}${monetary}${p}${fraction}`;
}

