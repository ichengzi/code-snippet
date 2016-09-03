private Color getColorFromName(string colorName)
{
    try
    {
        var colorObj = ColorConverter.ConvertFromString(colorName);
        var color = colorObj == null ? Colors.Green : (Color)colorObj;
        return color;
    }
    catch (Exception)
    {
        return Colors.Green;   
    }
}