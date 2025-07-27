export function timeAgo(timestamp, locale = 'en') {
    let value;
    const diff = Math.floor((new Date().getTime() - new Date(timestamp).getTime()) / 1000);
    const minutes = Math.floor(diff / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);
    const months = Math.floor(days / 30);
    const years = Math.floor(months / 12);

    const rtf = new Intl.RelativeTimeFormat(locale, { numeric: "auto" });
    if (years > 0) {
        value = rtf.format(-  years, "year");
    } else if (months > 0) {
        value = rtf.format(-  months, "month");
    } else if (days > 0) {
        value = rtf.format(-  days, "day");
    } else if (hours > 0) {
        value = rtf.format(-  hours, "hour");
    } else if (minutes > 0) {
        value = rtf.format(-  minutes, "minute");
    } else {
        value = rtf.format(-  diff, "second");
    }
    return value;
}