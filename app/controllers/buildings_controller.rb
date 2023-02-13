class BuildingsController < ApplicationController
  
  def index
    @items = Building.all.limit(50)
  end

  def show
    @building = Building.find_by_guid(params[:id])
  end

end
